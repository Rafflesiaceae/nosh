package fs

import (
	"os"

	"github.com/Rafflesiaceae/nosh/internal"
	"go.starlark.net/starlark"
)

func exists(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var paths = make([]string, 0)
	var dir, file bool = true, true

	if paths, err = internal.UnpackPositionalVarargsString("exists", args); err != nil {
		return starlark.Bool(false), err
	}

	if err := internal.UnpackKwargs("exists", kwargs, "file?", &file, "dir?", &dir); err != nil {
		return starlark.Bool(false), err
	}

	for _, path := range paths {
		if fi, err := os.Stat(path); err != nil {
			return starlark.Bool(false), nil
		} else if dir && fi.IsDir() {
			return starlark.Bool(true), nil
		} else if file && !fi.IsDir() {
			return starlark.Bool(true), nil
		}
	}

	return starlark.Bool(false), nil
}
