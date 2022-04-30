package os

import (
	"bytes"
	"fmt"
	"io"
	"nosh/internal"
	"os"
	"os/exec"

	"go.starlark.net/starlark"
	strlk "go.starlark.net/starlark"
	"go.starlark.net/syntax"
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
func (rr *RunResult) String() string        { return fmt.Sprintf("%v (%d)", rr.Cmd, rr.ExitCode) }

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

	var runArgs []string
	var check starlark.Bool = true
	var capture = runCaptureDefault
	var debug = false

	var captureStderr, captureStdout bool
	var devNullStderr, devNullStdout bool

	// Unpack args
	for _, arg := range args {
		// name, skipNone := paramName(pairs[2*i])
		// if skipNone {
		// 	if _, isNone := arg.(NoneType); isNone {
		// 		continue
		// 	}
		// }
		var val string
		if err := internal.UnpackOneArg(arg, &val); err != nil {
			// return nil, fmt.Errorf("%s: for parameter %s: %s", fnname, name, err)
			return nil, fmt.Errorf("unpacking: %s", arg)
		}
		runArgs = append(runArgs, val)
	}

	if err := internal.UnpackKwargs("run", args, kwargs, "check?", &check, "capture?", &capture, "debug?", &debug); err != nil {
		return nil, err
	}

	{ // parse capture
		iter := capture.Iterate()
		defer iter.Done()

		var el starlark.Value
		for iter.Next(&el) {

			if cmp, err := starlark.Compare(syntax.EQL, el, runCaptureStderr); err == nil && cmp == true {
				captureStderr = true
			}
			if cmp, err := starlark.Compare(syntax.EQL, el, runCaptureStderrDevNull); err == nil && cmp == true {
				captureStderr = false
				devNullStderr = true
			}
			if cmp, err := starlark.Compare(syntax.EQL, el, runCaptureStdout); err == nil && cmp == true {
				captureStdout = true
			}
			if cmp, err := starlark.Compare(syntax.EQL, el, runCaptureStdoutDevNull); err == nil && cmp == true {
				captureStdout = false
				devNullStdout = true
			}
		}
	}

	// Setup command
	cmd := exec.Command(runArgs[0], runArgs[1:]...)

	var stdoutBuf, stderrBuf bytes.Buffer
	if captureStdout {
		cmd.Stdout = &stdoutBuf
	} else if devNullStdout {
		cmd.Stdout = io.Discard
	} else {
		cmd.Stdout = os.Stdout
	}
	if captureStderr {
		cmd.Stderr = &stderrBuf
	} else if devNullStderr {
		cmd.Stderr = io.Discard
	} else {
		cmd.Stderr = os.Stderr
	}

	// Run command
	if debug {
		fmt.Fprintf(os.Stderr, "+ %s\n", cmd.String())
	}
	err = cmd.Run()

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
		Stdout:   outStr,
		Stderr:   errStr,
		ExitCode: exitCode,
	}, nil
}
