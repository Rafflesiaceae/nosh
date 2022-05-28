package os

import (
	"fmt"
	"os"

	"go.starlark.net/starlark"
)

func Fail(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		msg      string
		throw    bool
		exitCode = 1
	)

	if err := starlark.UnpackArgs("fail", args, kwargs, "msg", &msg, "exit_code?", &exitCode, "throw?", &throw); err != nil {
		return nil, err
	}

	PresetExitCode = exitCode

	if throw {
		return starlark.None, fmt.Errorf("%s", msg)
	}

	fmt.Fprintf(os.Stderr, "%s\n", msg)
	PresetExit()

	return starlark.None, nil // @XXX noop
}
