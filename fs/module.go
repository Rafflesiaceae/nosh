package fs

import (
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var Module = &starlarkstruct.Module{
	Name: "fs",
	Members: starlark.StringDict{
		"find": starlark.NewBuiltin("fs.find", find),
	},
}
