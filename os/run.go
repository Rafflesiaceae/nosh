package os

import (
	"bytes"
	"fmt"
	"nosh/internal"
	"os"
	"os/exec"

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
	runCaptureStdout  starlark.String
	runCaptureStderr  starlark.String
)

func init() {
	var err error

	// init runCaptureStdout
	runCaptureStdout = starlark.String("stdout")

	// init runCaptureStderr
	runCaptureStderr = starlark.String("stderr")

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
	var captureStdout, captureStderr bool

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

	if err := internal.UnpackKwargs("run", args, kwargs, "check?", &check, "capture?", &capture); err != nil {
		return nil, err
	}

	if internal.Contains(capture, runCaptureStdout) {
		captureStdout = true
	}
	if internal.Contains(capture, runCaptureStderr) {
		captureStderr = true
	}

	// Setup command
	cmd := exec.Command(runArgs[0], runArgs[1:]...)

	var stdoutBuf, stderrBuf bytes.Buffer
	if captureStdout {
		cmd.Stdout = &stdoutBuf
	} else {
		cmd.Stdout = os.Stdout
	}
	if captureStderr {
		cmd.Stderr = &stderrBuf
	} else {
		cmd.Stderr = os.Stderr
	}

	// Run command
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
