package fs

import (
	"fmt"
	"os"
	"runtime"
	"syscall"

	"go.starlark.net/starlark"
)

func move(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		from  string
		to    string
		force bool = false
	)

	if err = starlark.UnpackArgs("move", args, kwargs, "from", &from, "to", &to, "force?", &force); err != nil {
		return nil, err
	}

	if !force {
		if _, err := os.Stat(to); !os.IsNotExist(err) {
			return nil, fmt.Errorf("to path: \"%s\" already exist", to)
		}
	}

	if err = os.Rename(from, to); err != nil {
		// Rename will fail between different physical devices, attempt to copy in such cases.
		switch errorType := err.(type) {
		case *os.LinkError:
			if runtime.GOOS == "windows" {
				// Windows wraps it's own error codes.
				// https://msdn.microsoft.com/en-us/library/cc231199.aspx

				if errNo, ok := errorType.Err.(syscall.Errno); ok && errNo == 0x11 /* ERROR_NOT_SAME_DEVICE */ {
					return starlark.None, movePathViaCopyAndRemove(from, to, force)
				}

			} else if errorType.Err == syscall.EXDEV {
				return starlark.None, movePathViaCopyAndRemove(from, to, force)
			}
		}

		return nil, err
	}

	return starlark.None, nil
}

func movePathViaCopyAndRemove(from string, to string, force bool) error {
	if err := CopyPath(from, to, force); err != nil {
		return err
	}

	return os.RemoveAll(from)
}
