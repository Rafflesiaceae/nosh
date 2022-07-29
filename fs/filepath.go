package fs

import (
	"path/filepath"

	"github.com/Rafflesiaceae/nosh/internal"
	"go.starlark.net/starlark"
)

func abs(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		path   string
		result string
	)

	if err = starlark.UnpackArgs("abs", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	if result, err = filepath.Abs(path); err != nil {
		return nil, err
	}

	return starlark.String(result), nil
}

func base(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var path string

	if err = starlark.UnpackArgs("base", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	return starlark.String(filepath.Base(path)), nil
}

func clean(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var path string

	if err = starlark.UnpackArgs("clean", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	return starlark.String(filepath.Clean(path)), nil
}

func dir(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var path string

	if err = starlark.UnpackArgs("dir", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	return starlark.String(filepath.Dir(path)), nil
}

func ext(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var path string

	if err = starlark.UnpackArgs("ext", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	return starlark.String(filepath.Ext(path)), nil
}

func join(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var paths = make([]string, 0)

	if paths, err = internal.UnpackPositionalVarargsString("join", args); err != nil {
		return nil, err
	}

	return starlark.String(filepath.Join(paths...)), nil
}

func rel(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		result   string
		basepath string
		targpath string
	)

	if err = starlark.UnpackArgs("rel", args, kwargs, "basepath", &basepath, "targpath", &targpath); err != nil {
		return nil, err
	}

	if result, err = filepath.Rel(basepath, targpath); err != nil {
		return nil, err
	}

	return starlark.String(result), nil
}

func split(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		resultDir  string
		resultFile string
		path       string
	)

	if err = starlark.UnpackArgs("split", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	resultDir, resultFile = filepath.Split(path)

	return starlark.NewList(
			[]starlark.Value{
				starlark.String(resultDir),
				starlark.String(resultFile),
			}),
		nil
}

func splitList(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var path string

	if err = starlark.UnpackArgs("split_list", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	result := make([]starlark.Value, 0)
	for _, part := range filepath.SplitList(path) {
		result = append(result, starlark.String(part))
	}

	return starlark.NewList(result), nil
}

func toSlash(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var path string

	if err = starlark.UnpackArgs("to_slash", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	return starlark.String(filepath.ToSlash(path)), nil
}
