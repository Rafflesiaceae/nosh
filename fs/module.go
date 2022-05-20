package fs

import (
	"os"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var ModuleChdir = starlark.NewBuiltin("fs.chdir", chdir)
var ModuleCopy = starlark.NewBuiltin("fs.copy", copyImpl)
var ModuleExists = starlark.NewBuiltin("fs.exists", exists)
var ModuleFind = starlark.NewBuiltin("fs.find", find)
var ModuleMkdir = starlark.NewBuiltin("fs.mkdir", mkdir)
var ModuleMove = starlark.NewBuiltin("fs.move", move)
var ModulePwd = starlark.NewBuiltin("fs.pwd", pwd)
var ModuleRead = starlark.NewBuiltin("fs.read", read)
var ModuleRemove = starlark.NewBuiltin("fs.remove", remove)
var ModuleTouch = starlark.NewBuiltin("fs.touch", touch)
var ModuleWrite = starlark.NewBuiltin("fs.write", write)

var Module *starlarkstruct.Module

func init() {
	Module = &starlarkstruct.Module{
		Name: "fs",
		Members: starlark.StringDict{
			"chdir":               ModuleChdir,
			"copy":                ModuleCopy,
			"exists":              ModuleExists,
			"find":                ModuleFind,
			"mkdir":               ModuleMkdir,
			"move":                ModuleMove,
			"path_list_separator": starlark.String(os.PathListSeparator),
			"path_separator":      starlark.String(os.PathSeparator),
			"pwd":                 ModulePwd,
			"read":                ModuleRead,
			"remove":              ModuleRemove,
			"touch":               ModuleTouch,
			"write":               ModuleWrite,
		},
	}

}
