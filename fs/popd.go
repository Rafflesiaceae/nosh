package fs

import (
	"os"

	"go.starlark.net/starlark"
)

func popd(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		reset bool = false
	)
	if err = starlark.UnpackArgs("popd", args, kwargs, "reset?", &reset); err != nil {
		return nil, err
	}

	if reset {
		dirStack.Clear()
		return starlark.None, nil
	}

	if dirStack.Len() == 0 {
		return starlark.None, nil
	}

	last := popDirStack()
	if err := os.Chdir(last.GoString()); err != nil {
		return nil, err
	}

	return last, nil
}
