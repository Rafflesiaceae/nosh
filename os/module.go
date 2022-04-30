package os

import (
	"os"
	"time"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var ModuleExit = starlark.NewBuiltin("os.exit", quit)
var ModuleQuit = starlark.NewBuiltin("os.quit", quit)
var ModuleRun = starlark.NewBuiltin("os.run", run)

var Module = &starlarkstruct.Module{
	Name: "os",
	Members: starlark.StringDict{
		"distro":     starlark.NewBuiltin("os.distro", distro),
		"executable": starlark.NewBuiltin("os.executable", executable),
		"exit":       ModuleExit,
		"quit":       ModuleQuit,
		"run":        ModuleRun,
		"sleep":      starlark.NewBuiltin("os.sleep", sleep),
	},
}

func executable(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	if err := starlark.UnpackArgs("executable", args, kwargs); err != nil {
		return nil, err
	}

	result, err := os.Executable()
	if err != nil {
		return nil, err
	}

	return starlark.String(result), nil
}

func quit(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	exitCode, err := starlark.AsInt32(args.Index(0))
	if err != nil {
		panic(err)
	}

	os.Exit(exitCode)
	return starlark.None, nil // @XXX noop
}

func sleep(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	sleepTime, err := starlark.AsInt32(args.Index(0))
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	return starlark.None, nil
}
