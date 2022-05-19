package fs

import (
	"fmt"
	"os"

	"go.starlark.net/starlark"
)

func write(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		append   starlark.Bool = false
		path     string
		contents string
		force    bool = true
	)

	if err := starlark.UnpackArgs("write", args, kwargs, "path", &path, "contents", &contents, "append?", &append, "force?", &force); err != nil {
		return nil, err
	}

	if !force {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return nil, fmt.Errorf("path: \"%s\" already exist", path)
		}
	}

	flags := os.O_RDWR | os.O_CREATE
	if append {
		flags = flags | os.O_APPEND
	} else {
		flags = flags | os.O_TRUNC
	}

	f, err := os.OpenFile(path, flags, 0755)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = f.Write([]byte(contents))
	if err != nil {
		return nil, err
	}

	return starlark.None, nil
}
