package lang

import "go.starlark.net/starlark"

var (
	DeferStack = make([]starlark.Callable, 0)
)

func Defer(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		callable starlark.Callable
	)

	if err := starlark.UnpackArgs("defer", args, kwargs, "callable", &callable); err != nil {
		return nil, err
	}

	DeferStack = append(DeferStack, callable)

	return starlark.None, nil
}
