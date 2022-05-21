package fs

import (
	"go.starlark.net/starlark"
)

func cmpPath(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		firstPath  string
		secondPath string
	)

	if err = starlark.UnpackArgs("cmp_file", args, kwargs, "first_path", &firstPath, "second_path", &secondPath); err != nil {
		return nil, err
	}

	return starlark.Bool(IsSamePath(firstPath, secondPath)), nil
}
