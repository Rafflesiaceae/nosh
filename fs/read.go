package fs

import (
	"io"
	"os"

	"go.starlark.net/starlark"
)

func read(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		path = starlark.String("")
	)

	if err = starlark.UnpackArgs("read", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	switch path {
	case "<stdin>":
		bytes, err := io.ReadAll(os.Stdin)
		if err == nil {

		}

		return starlark.String(bytes), nil
	default:
		contents, err := os.ReadFile(path.GoString())
		if err != nil {
			return nil, err
		}

		return starlark.String(contents), nil
	}
}
