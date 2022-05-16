package fs

import (
	"io/ioutil"

	"go.starlark.net/starlark"
	strlk "go.starlark.net/starlark"
)

func read(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		path = strlk.String("")
	)

	if err = starlark.UnpackPositionalArgs("read", args, kwargs, 1, &path); err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadFile(path.GoString())
	if err != nil {
		return nil, err
	}

	return starlark.String(contents), nil
}