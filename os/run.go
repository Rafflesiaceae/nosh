package os

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/Rafflesiaceae/nosh/internal"

	"go.starlark.net/starlark"
	strlk "go.starlark.net/starlark"
)

type RunResult struct { // {{{
	Cmd      []string
	Stdout   string
	Stderr   string
	ExitCode int
}

var (
	_ strlk.Value    = &RunResult{}
	_ strlk.HasAttrs = &RunResult{}
)

func (rr *RunResult) Freeze()               { /* @TODO */ }
func (rr *RunResult) Hash() (uint32, error) { /* @TODO */ return 0, nil }
func (rr *RunResult) Truth() strlk.Bool     { return true }
func (rr *RunResult) Type() string          { return "RunResult" }
func (rr *RunResult) String() string {

	// build capture summary
	summary := make([]string, 0)
	summary = append(summary, fmt.Sprintf("%d", rr.ExitCode))
	if rr.Stdout != "" {
		summary = append(summary, "<stdout>")
	}
	if rr.Stderr != "" {
		summary = append(summary, "<stderr>")
	}

	// build arguments
	args := make([]string, 0)
	for _, arg := range rr.Cmd {
		res := arg
		if strings.ContainsAny(arg, " (){}[]*$") {
			// add single-quotes around args if they contain special chars
			res = fmt.Sprintf("'%s'", arg)
		}

		args = append(args, res)
	}

	return fmt.Sprintf(
		"%s (%s)",
		strings.Join(args, " "),
		strings.Join(summary, ", "),
	)
}

func (rr *RunResult) Attr(name string) (strlk.Value, error) {
	return map[string]strlk.Value{
		// @TODO Cmd
		"stdout":   strlk.String(rr.Stdout),
		"stderr":   strlk.String(rr.Stderr),
		"exitCode": strlk.MakeInt(rr.ExitCode),
	}[name], nil
}

func (fe *RunResult) AttrNames() []string {
	return []string{"stdout", "stderr", "exitCode"}
}

// }}}

var (
	runCaptureDefault *starlark.List
	runCaptureStderr  starlark.String
	runCaptureStdout  starlark.String
	// @XXX @TODO capture != redirection, this needs to be a different concept
	runCaptureStdoutDevNull starlark.String
	runCaptureStderrDevNull starlark.String
)

func init() {
	var err error

	runCaptureStderr = starlark.String("stderr")
	runCaptureStderrDevNull = starlark.String("stderr->devnull")
	runCaptureStdout = starlark.String("stdout")
	runCaptureStdoutDevNull = starlark.String("stdout->devnull")

	// init runCaptureDefault
	runCaptureDefault = &starlark.List{}

	err = runCaptureDefault.Append(runCaptureStdout)
	if err != nil {
		panic(err)
	}
	err = runCaptureDefault.Append(runCaptureStderr)
	if err != nil {
		panic(err)
	}
}

func run(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		runArgs []string
		check   starlark.Bool = true
		capture               = &starlark.List{}
		debug                 = false
		env     *starlark.List
	)

	var captureStderr, captureStdout bool
	var appendStderr, appendStdout bool
	var redirStderr, redirStdout string
	var stdin string

	runArgs, err = internal.UnpackPositionalVarargsString("run", args)
	if err != nil {
		return nil, err
	}

	if err := starlark.UnpackArgs("run", nil, kwargs, "check?", &check, "capture?", &capture, "debug?", &debug, "env?", &env); err != nil {
		return nil, err
	}

	{ // parse capture
		iter := capture.Iterate()
		defer iter.Done()

		var el starlark.Value
		for iter.Next(&el) {

			var str string
			if starStr, ok := el.(starlark.String); ok {
				str = starStr.GoString()
			} else {
				return nil, fmt.Errorf("expected string, got: %v", el)
			}

			switch str {
			case "stdout":
				captureStdout = true
				continue
			case "stderr":
				captureStderr = true
				continue
			}

			if strings.HasPrefix(str, "stdout->>") {
				redirStdout = strings.SplitN(str, "stdout->>", 2)[1]
				appendStdout = true
				continue
			} else if strings.HasPrefix(str, "stderr->>") {
				redirStderr = strings.SplitN(str, "stderr->>", 2)[1]
				appendStderr = true
				continue
			} else if strings.HasPrefix(str, "stdout->") {
				redirStdout = strings.SplitN(str, "stdout->", 2)[1]
				continue
			} else if strings.HasPrefix(str, "stderr->") {
				redirStderr = strings.SplitN(str, "stderr->", 2)[1]
				continue
			} else if strings.HasPrefix(str, "stdin<-") {
				stdin = strings.SplitN(str, "stdin<-", 2)[1]
				continue
			}

			return nil, fmt.Errorf("unsupported capture: '%s'", el)
		}
	}

	// Setup command
	cmd := exec.Command(runArgs[0], runArgs[1:]...)

	var stdoutBuf, stderrBuf bytes.Buffer
	if captureStdout {
		cmd.Stdout = &stdoutBuf
	} else if redirStdout != "" && redirStdout == "devnull" {
		cmd.Stdout = io.Discard
	} else if redirStdout != "" {
		var f *os.File
		flags := os.O_RDWR | os.O_CREATE
		if appendStdout {
			flags = flags | os.O_APPEND
		} else {
			flags = flags | os.O_TRUNC
		}
		f, err = os.OpenFile(redirStdout, flags, 0755)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		cmd.Stdout = f
	} else {
		cmd.Stdout = os.Stdout
	}
	if captureStderr {
		cmd.Stderr = &stderrBuf
	} else if redirStderr != "" && redirStderr == "devnull" {
		cmd.Stderr = io.Discard
	} else if redirStderr != "" {
		var f *os.File
		flags := os.O_RDWR | os.O_CREATE
		if appendStderr {
			flags = flags | os.O_APPEND
		} else {
			flags = flags | os.O_TRUNC
		}
		f, err = os.OpenFile(redirStderr, flags, 0755)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		cmd.Stderr = f
	} else {
		cmd.Stderr = os.Stderr
	}

	var stdinFile *os.File
	if stdin == "" {
		goto postStdinAssignment
	}

	stdinFile, err = os.Open(stdin)
	if err != nil {
		return nil, err
	}
	defer stdinFile.Close()

	cmd.Stdin = stdinFile

postStdinAssignment:
	if env != nil {
		result := make([]string, 0)
		result = append(result, os.Environ()...)

		iter := env.Iterate()
		defer iter.Done()

		var el starlark.Value
		for iter.Next(&el) {
			if str, ok := el.(starlark.String); ok {
				result = append(result, str.GoString())
			}
		}

		cmd.Env = result
	}

	// Run command
	if debug {
		fmt.Fprintf(os.Stderr, "+ %s\n", cmd.String())
	}
	err = cmd.Run()
	// @TODO do I need to flush on Windows here?

	// Gather command results
	var exitCode = 0
	if exitError, ok := err.(*exec.ExitError); ok {
		exitCode = exitError.ExitCode()
	}

	if check && exitCode != 0 {
		return nil, fmt.Errorf("command exited with code %d", exitCode)
	}

	var outStr, errStr string
	if captureStdout {
		outStr = string(stdoutBuf.Bytes())
	}
	if captureStderr {
		errStr = string(stderrBuf.Bytes())
	}

	return &RunResult{
		Cmd:      runArgs,
		Stdout:   outStr,
		Stderr:   errStr,
		ExitCode: exitCode,
	}, nil
}
