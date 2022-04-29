package os

import (
	"time"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var Module = &starlarkstruct.Module{
	Name: "os",
	Members: starlark.StringDict{
		"sleep": starlark.NewBuiltin("os.sleep", sleep),
	},
}

func sleep(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	sleepTime, err := starlark.AsInt32(args.Index(0))
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	return starlark.None, nil
}
