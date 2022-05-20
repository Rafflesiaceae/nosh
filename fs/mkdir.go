package fs

import (
	"os"

	"github.com/Rafflesiaceae/nosh/internal"
	"go.starlark.net/starlark"
)

func mkdir(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var paths = make([]string, 0)

	if paths, err = internal.UnpackPositionalVarargsString("mkdir", args); err != nil {
		return nil, err
	}

	result := make([]starlark.Value, 0)
	for _, path := range paths {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return nil, err
		}
		result = append(result, starlark.String(path))
	}

	return starlark.NewList(result), nil
}
