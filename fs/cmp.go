package fs

import (
	"bytes"
	"io/ioutil"

	"go.starlark.net/starlark"
)

func cmp(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		firstPath  string
		secondPath string
	)

	if err = starlark.UnpackArgs("cmp", args, kwargs, "first_path", &firstPath, "second_path", &secondPath); err != nil {
		return nil, err
	}

	// @TODO impl dir comparing

	// @TODO optimize w/o reading everything into RAM
	var (
		firstContents  []byte
		secondContents []byte
	)
	if firstContents, err = ioutil.ReadFile(firstPath); err != nil {
		return nil, err
	}
	if secondContents, err = ioutil.ReadFile(secondPath); err != nil {
		return nil, err
	}

	return starlark.Bool(bytes.Equal(firstContents, secondContents)), nil
}
