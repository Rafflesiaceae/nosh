package os

import (
	"os"

	"github.com/Rafflesiaceae/nosh/lang"
	"go.starlark.net/starlark"
)

var (
	PresetExitCode = 0
)

func PresetExit(thread *starlark.Thread) {
	var err error

	// workf off DeferStack
	//
	// this currently happens after 'thread.ExecFile' has run, which freezes all
	// collections at the end ( @TODO determine if freezing at the end is
	// desirable, can be disabled simply by only calling the portions of
	// 'thread.ExecFile' until the freeze, might make sense to be configurable )
	for len(lang.DeferStack) > 0 {
		callable := lang.DeferStack[len(lang.DeferStack)-1]
		_, err = starlark.Call(thread, callable, starlark.Tuple{}, nil)
		if err != nil {
			panic(err)
		}

		lang.DeferStack = lang.DeferStack[:len(lang.DeferStack)-1]
	}

	os.Exit(PresetExitCode)
}

func exit(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	exitCode, err := starlark.AsInt32(args.Index(0))
	if err != nil {
		panic(err)
	}

	PresetExitCode = exitCode
	PresetExit(thread)
	return starlark.None, nil // @XXX noop
}
