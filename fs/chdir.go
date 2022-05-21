package fs

import (
	"os"

	"github.com/Rafflesiaceae/nosh/internal"
	"go.starlark.net/starlark"
)

func chdir(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var paths []string
	if paths, err = internal.UnpackPositionalVarargsString("chdir", args); err != nil {
		return starlark.Bool(false), err
	}

	pwd, _ := os.Getwd()
	for _, path := range paths {
		if err = os.Chdir(path); err != nil {
			return nil, err
		}
	}

	if pwd != "" && pwd != peekDirStack().GoString() {
		// only push previous working dir to dirStack if it is not the same as
		// the last path stored in dirStack

		dirStack.Append(starlark.String(pwd))
	}

	return starlark.String(pwd), nil
}
