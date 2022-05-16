package fs

import (
	"os"
	"path"

	"github.com/Rafflesiaceae/nosh/internal"
	"go.starlark.net/starlark"
)

func touch(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var paths []string
	if paths, err = internal.UnpackPositionalVarargsString("touch", args); err != nil {
		return nil, err
	}

	for _, pth := range paths {
		if _, err := os.Stat(pth); err == nil {
			continue
		}

		dir, file := path.Split(pth)

		err = os.MkdirAll(dir, 0700)
		if err != nil {
			return nil, err
		}

		if file != "" {
			f, err := os.Create(pth)
			if err != nil {
				return nil, err
			}
			f.Close()
		}
	}

	return starlark.None, nil
}
