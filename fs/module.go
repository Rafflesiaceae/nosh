package fs

import (
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var ModuleExists = starlark.NewBuiltin("fs.exists", exists)
var ModuleFind = starlark.NewBuiltin("fs.find", find)
var ModuleMove = starlark.NewBuiltin("fs.move", move)
var ModuleRead = starlark.NewBuiltin("fs.read", read)
var ModuleRemove = starlark.NewBuiltin("fs.remove", remove)
var ModuleTouch = starlark.NewBuiltin("fs.touch", touch)
var ModuleWrite = starlark.NewBuiltin("fs.write", write)

var Module = &starlarkstruct.Module{
	Name: "fs",
	Members: starlark.StringDict{
		"exists": ModuleExists,
		"find":   ModuleFind,
		"move":   ModuleMove,
		"read":   ModuleRead,
		"remove": ModuleRemove,
		"touch":  ModuleTouch,
		"write":  ModuleWrite,
	},
}
