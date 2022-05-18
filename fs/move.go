package fs

import (
	"os"
	"runtime"
	"syscall"

	"go.starlark.net/starlark"
)

func move(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		from string
		to   string
	)

	if err = starlark.UnpackPositionalArgs("move", args, kwargs, 2, &from, &to); err != nil {
		return nil, err
	}

	if err = os.Rename(from, to); err != nil {
		// Rename will fail between different physical devices, attempt to copy in such cases.
		switch errorType := err.(type) {
		case *os.LinkError:
			if runtime.GOOS == "windows" {
				// Windows wraps it's own error codes.
				// https://msdn.microsoft.com/en-us/library/cc231199.aspx

				if errNo, ok := errorType.Err.(syscall.Errno); ok && errNo == 0x11 /* ERROR_NOT_SAME_DEVICE */ {
					return starlark.None, MovePath(from, to)
				}

			} else if errorType.Err == syscall.EXDEV {
				return starlark.None, MovePath(from, to)
			}
		}

		return nil, err
	}

	return starlark.None, nil
}
