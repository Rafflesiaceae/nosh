package fs

import (
	"path"

	"github.com/Rafflesiaceae/nosh/internal"
	"go.starlark.net/starlark"
)

func join(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var paths = make([]string, 0)

	if paths, err = internal.UnpackPositionalVarargsString("join", args); err != nil {
		return starlark.Bool(false), err
	}

	return starlark.String(path.Join(paths...)), nil
}
