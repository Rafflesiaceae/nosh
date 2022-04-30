package os

import (
	"os"
	"time"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var ModuleExit = starlark.NewBuiltin("os.exit", quit)
var ModuleExpand = starlark.NewBuiltin("os.expand", expand)
var ModuleGetenv = starlark.NewBuiltin("os.getenv", getenv)
var ModuleQuit = starlark.NewBuiltin("os.quit", quit)
var ModuleRun = starlark.NewBuiltin("os.run", run)
var ModuleSetenv = starlark.NewBuiltin("os.setenv", setenv)

var Module = &starlarkstruct.Module{
	Name: "os",
	Members: starlark.StringDict{
		"distro":     starlark.NewBuiltin("os.distro", distro),
		"executable": starlark.NewBuiltin("os.executable", executable),
		"exit":       ModuleExit,
		"expand":     ModuleExpand,
		"getenv":     ModuleGetenv,
		"quit":       ModuleQuit,
		"run":        ModuleRun,
		"setenv":     ModuleSetenv,
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

func expand(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var str starlark.String
	if err := starlark.UnpackPositionalArgs("expand", args, kwargs, 1, &str); err != nil {
		return nil, err
	}

	return starlark.String(os.ExpandEnv(str.GoString())), nil
}

func getenv(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var key starlark.String
	if err := starlark.UnpackPositionalArgs("getenv", args, kwargs, 0, &key); err != nil {
		return nil, err
	}

	if key != "" { // return single env val
		return starlark.String(os.Getenv(key.GoString())), nil
	} else { // return whole env
		environ := os.Environ()

		listContent := make([]starlark.Value, 0)
		for _, v := range environ {
			listContent = append(listContent, starlark.String(v))
		}

		return starlark.NewList(listContent), nil
	}
}

func quit(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	exitCode, err := starlark.AsInt32(args.Index(0))
	if err != nil {
		panic(err)
	}

	os.Exit(exitCode)
	return starlark.None, nil // @XXX noop
}

func setenv(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var key, val starlark.String
	if err := starlark.UnpackPositionalArgs("setenv", args, kwargs, 2, &key, &val); err != nil {
		return nil, err
	}

	os.Setenv(key.GoString(), val.GoString())

	return starlark.None, nil
}

func sleep(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	sleepTime, err := starlark.AsInt32(args.Index(0))
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	return starlark.None, nil
}
