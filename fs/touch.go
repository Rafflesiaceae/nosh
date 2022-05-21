package fs

import (
	"os"
	"path/filepath"

	"github.com/Rafflesiaceae/nosh/internal"
	"go.starlark.net/starlark"
)

func touch(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var paths []string
	if paths, err = internal.UnpackPositionalVarargsString("touch", args); err != nil {
		return nil, err
	}

	for _, path := range paths {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			continue
		}

		dir := filepath.Dir(path)

		err = os.MkdirAll(dir, 0700)
		if err != nil {
			return nil, err
		}

		f, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		f.Close()
	}

	return starlark.None, nil
}
