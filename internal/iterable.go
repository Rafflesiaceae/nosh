package internal

import (
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

// Contains checks if a starlark.Iterable contains an element
func Contains(iterable starlark.Iterable, target starlark.Value) bool {
	iter := iterable.Iterate()
	defer iter.Done()

	var el starlark.Value
	for iter.Next(&el) {

		cmp, err := starlark.Compare(syntax.EQL, target, el)
		if err == nil && cmp == true {
			return true
		}

	}
	return false
}
