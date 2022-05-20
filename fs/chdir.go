package fs

import (
	"os"

	"go.starlark.net/starlark"
)

func chdir(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		path string
	)

	if err := starlark.UnpackArgs("chdir", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	pwd, _ := os.Getwd()

	if err := os.Chdir(path); err != nil {
		return nil, err
	}

	if pwd != "" {
		dirStack.Append(starlark.String(pwd))
	}

	return starlark.String(path), nil
}
