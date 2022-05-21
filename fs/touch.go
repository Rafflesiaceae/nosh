package fs

import (
	"fmt"
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

		pathLastChar := path[len(path)-1]
		if pathLastChar == os.PathSeparator || pathLastChar == '/' {
			return nil, fmt.Errorf("can't touch a directory path: \"%s\"", path)
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
