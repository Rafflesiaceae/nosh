package os

import (
	"go.starlark.net/starlark"
)

// Args are the arguments accessible by nosh
var Args *starlark.List = starlark.NewList([]starlark.Value{})

// SetArgs initializes the args accessible by nosh
func SetArgs(args []string) {
	result := make([]starlark.Value, 0)
	for _, v := range args {
		result = append(result, starlark.String(v))
	}
	Args = starlark.NewList(result)
}
