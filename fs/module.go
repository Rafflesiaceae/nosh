package fs

import (
	"os"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var ModuleChdir = starlark.NewBuiltin("fs.chdir", chdir)
var ModuleCmp = starlark.NewBuiltin("fs.cmp", cmp)
var ModuleCmpFile = starlark.NewBuiltin("fs.cmp_file", cmpFile)
var ModuleCmpPath = starlark.NewBuiltin("fs.cmp_path", cmpPath)
var ModuleCopy = starlark.NewBuiltin("fs.copy", copyImpl)
var ModuleExists = starlark.NewBuiltin("fs.exists", exists)
var ModuleFind = starlark.NewBuiltin("fs.find", find)
var ModuleJoin = starlark.NewBuiltin("fs.join", join)
var ModuleMkdir = starlark.NewBuiltin("fs.mkdir", mkdir)
var ModuleMove = starlark.NewBuiltin("fs.move", move)
var ModulePopd = starlark.NewBuiltin("fs.popd", popd)
var ModulePwd = starlark.NewBuiltin("fs.pwd", pwd)
var ModuleRead = starlark.NewBuiltin("fs.read", read)
var ModuleReaddir = starlark.NewBuiltin("fs.readdir", readdir)
var ModuleRemove = starlark.NewBuiltin("fs.remove", remove)
var ModuleTouch = starlark.NewBuiltin("fs.touch", touch)
var ModuleWatch = starlark.NewBuiltin("fs.watch", watch)
var ModuleWrite = starlark.NewBuiltin("fs.write", write)

var Module *starlarkstruct.Module

func init() {
	Module = &starlarkstruct.Module{
		Name: "fs",
		Members: starlark.StringDict{
			"abs":                 starlark.NewBuiltin("fs.abs", abs),
			"base":                starlark.NewBuiltin("fs.base", base),
			"chdir":               ModuleChdir,
			"clean":               starlark.NewBuiltin("fs.clean", clean),
			"cmp":                 ModuleCmp,
			"cmp_file":            ModuleCmpFile,
			"cmp_path":            ModuleCmpPath,
			"copy":                ModuleCopy,
			"dir":                 starlark.NewBuiltin("fs.dir", dir),
			"dir_stack":           dirStack,
			"exists":              ModuleExists,
			"ext":                 starlark.NewBuiltin("fs.ext", ext),
			"find":                ModuleFind,
			"join":                ModuleJoin,
			"mkdir":               ModuleMkdir,
			"move":                ModuleMove,
			"parse":               starlark.NewBuiltin("fs.parse", parse),
			"path_list_separator": starlark.String(os.PathListSeparator),
			"path_separator":      starlark.String(os.PathSeparator),
			"popd":                ModulePopd,
			"pwd":                 ModulePwd,
			"read":                ModuleRead,
			"readdir":             ModuleReaddir,
			"rel":                 starlark.NewBuiltin("fs.rel", rel),
			"remove":              ModuleRemove,
			"split":               starlark.NewBuiltin("fs.split", split),
			"split_list":          starlark.NewBuiltin("fs.split_list", splitList),
			"to_slash":            starlark.NewBuiltin("fs.to_slash", toSlash),
			"touch":               ModuleTouch,
			"watch":               ModuleWatch,
			"write":               ModuleWrite,
		},
	}

}
