package fs

import (
	"os"
	"path/filepath"

	"go.starlark.net/starlark"
)

func readdir(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		path string
	)

	if err = starlark.UnpackArgs("readdir", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := make([]starlark.Value, 0)
	for _, file := range files {
		ne, err := FromPath(filepath.Join(absPath, file.Name()))
		if err != nil {
			continue
		}

		result = append(result, ne)
	}

	return starlark.NewList(result), nil
}
