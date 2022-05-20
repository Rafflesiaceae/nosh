package fs

import (
	"os"

	"go.starlark.net/starlark"
)

func pwd(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	if err := starlark.UnpackArgs("pwd", args, kwargs); err != nil {
		return nil, err
	}

	var pwd string
	if pwd, err = os.Getwd(); err != nil {
		return nil, err
	}

	return starlark.String(pwd), nil
}
