package fs

import (
	"os"

	"go.starlark.net/starlark"
)

func move(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		from string
		to   string
	)

	if err = starlark.UnpackPositionalArgs("move", args, kwargs, 2, &from, &to); err != nil {
		return nil, err
	}

	if err = os.Rename(from, to); err != nil {
		return nil, err
	}

	return starlark.None, nil
}
