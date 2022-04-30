package lang

import (
	"fmt"

	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

func Assert(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	// @TODO return an AssertionError that also involves a backtrace
	var err error

	var x, y starlark.Comparable
	var xfail starlark.Bool

	if err := starlark.UnpackArgs("assert", args, kwargs, "x", &x, "y?", &y, "xfail?", &xfail); err != nil {
		return nil, err
	}

	if y == nil {
		y = starlark.True
	}

	truth, err := x.CompareSameType(syntax.EQL, y, 1)
	if err != nil {
		return nil, err
	}

	if xfail {
		truth = !truth
	}

	if !truth {
		// println("ASDWWEE")
		return starlark.None, fmt.Errorf("%s != %s", x, y)
	}

	return starlark.None, nil
}
