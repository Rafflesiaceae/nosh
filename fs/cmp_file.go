package fs

import (
	"os"

	"go.starlark.net/starlark"
)

func cmpFile(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		firstPath  string
		secondPath string
	)

	if err = starlark.UnpackArgs("cmp_file", args, kwargs, "first_path", &firstPath, "second_path", &secondPath); err != nil {
		return nil, err
	}

	var firstFi os.FileInfo
	var secondFi os.FileInfo
	if firstFi, err = os.Stat(firstPath); err != nil {
		return nil, err
	}
	if secondFi, err = os.Stat(firstPath); err != nil {
		return nil, err
	}

	return starlark.Bool(os.SameFile(firstFi, secondFi)), nil
}
