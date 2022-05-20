package fs

import "go.starlark.net/starlark"

var dirStack = starlark.NewList([]starlark.Value{})

func popDirStack() starlark.String {
	result := dirStack.Index(dirStack.Len() - 1).(starlark.String)
	updateDirStack(dirStack.Slice(0, dirStack.Len()-1, 1).(*starlark.List)) // pop
	return result
}

func updateDirStack(val *starlark.List) {
	dirStack = val
	Module.Members["dir_stack"] = val
}
