package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

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

	var pwd string
	gatheredPwd := false
	gatherPwd := func() {
		if pwd, err = os.Getwd(); err != nil {
			pwd = ""
		}
	}

	for _, path := range paths {
		if !force { // fail if path doesn't exist w/o force
			if _, err := os.Stat(path); os.IsNotExist(err) {
				return nil, fmt.Errorf("path: \"%s\" doesn't exist", path)
			}
		}

		if debug {
			fmt.Fprintf(os.Stderr, "+ rm -rf? %s", path)
		}

		if runtime.GOOS == "windows" {
			// on windows we can't remove a directory we are currently in due to
			// mandatory locking
			//
			// we check if our current workdir is within the path to be removed
			// and if try to move out of the way before attempting to remove it

			if fi, err := os.Stat(path); err != nil {
				continue
			} else if fi.IsDir() {

				if !gatheredPwd {
					gatheredPwd = true
					gatherPwd()
				}

				pathWithin, err := IsPathWithin(path, pwd)
				if err != nil {
					return nil, err
				}

				if pathWithin {
					_, err := chdir(thread, fn, starlark.Tuple{starlark.String(filepath.Dir(path))}, nil)
					if err != nil {
						return nil, err
					}
				}
			}
		}

		err = os.RemoveAll(path)
		if err != nil {
			return nil, err
		}
	}

	return starlark.None, nil
}
