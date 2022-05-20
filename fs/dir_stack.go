package fs

import "go.starlark.net/starlark"

var dirStack = starlark.NewList([]starlark.Value{})

func peekDirStack() starlark.String {
	dirStackLen := dirStack.Len()
	if dirStackLen > 0 {
		return dirStack.Index(dirStack.Len() - 1).(starlark.String)
	}
	return starlark.String("")
}

func popDirStack() starlark.String {
	result := peekDirStack()
	updateDirStack(dirStack.Slice(0, dirStack.Len()-1, 1).(*starlark.List)) // pop
	return result
}

func updateDirStack(val *starlark.List) {
	dirStack = val
	Module.Members["dir_stack"] = val
}
