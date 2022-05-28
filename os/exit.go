package os

import (
	"os"

	"go.starlark.net/starlark"
)

var (
	PresetExitCode = 1
)

func PresetExit() {

	os.Exit(PresetExitCode)
}

func exit(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	exitCode, err := starlark.AsInt32(args.Index(0))
	if err != nil {
		panic(err)
	}

	PresetExitCode = exitCode
	PresetExit()
	return starlark.None, nil // @XXX noop
}
