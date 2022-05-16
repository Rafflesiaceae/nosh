package fs

import (
	"fmt"
	"os"

	"github.com/Rafflesiaceae/nosh/internal"
	"go.starlark.net/starlark"
)

func remove(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	// @TODO file, dir

	var paths []string
	if paths, err = internal.UnpackPositionalVarargsString("remove", args); err != nil {
		return nil, err
	}

	var force, debug bool = false, false

	if err = internal.UnpackKwargs("remove", kwargs, "force", &force, "debug", &debug); err != nil {
		return nil, err
	}

	for _, pth := range paths {
		if !force {
			if _, err := os.Stat(pth); err != nil {
				return nil, err
			}
		}

		if debug {
			fmt.Fprintf(os.Stderr, "+ rm -rf %s", pth)
		}

		err = os.RemoveAll(pth)
		if err != nil {
			return nil, err
		}
	}

	return starlark.None, nil
}
