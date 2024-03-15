package lang

import (
	"fmt"

	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

func Assert(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	// @TODO return an AssertionError that also involves a backtrace
	var err error

	var (
		msg   string
		x, y  starlark.Value
		xfail starlark.Bool
	)

	if err := starlark.UnpackArgs("assert", args, kwargs, "x", &x, "y?", &y, "msg?", &msg, "xfail?", &xfail); err != nil {
		return nil, err
	}

	if y == nil {
		x = x.Truth()
		y = starlark.True
	}

	truth, err := starlark.Compare(syntax.EQL, x, y)
	if err != nil {
		return nil, err
	}

	if xfail {
		truth = !truth
	}

	if !truth {
		if msg == "" {
			return starlark.None, fmt.Errorf("%s != %s", x, y)
		} else {
			return starlark.None, fmt.Errorf("%s; %s != %s", msg, x, y)
		}
	}

	return starlark.None, nil
}
