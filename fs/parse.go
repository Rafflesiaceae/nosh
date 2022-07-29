package fs

import (
	"fmt"
	"path/filepath"

	"go.starlark.net/starlark"
	strlk "go.starlark.net/starlark"
)

// FileEntry
type FileEntry struct {
	AbsPath  string
	BaseName string
	Dir      string
	Ext      string
	Name     string
	Path     string
}

var (
	_ strlk.Value    = &FileEntry{}
	_ strlk.HasAttrs = &FileEntry{}
)

func (fe *FileEntry) Freeze()               { /* @TODO */ }
func (fe *FileEntry) Hash() (uint32, error) { /* @TODO */ return 0, nil }
func (fe *FileEntry) Truth() strlk.Bool     { return true }
func (fe *FileEntry) Type() string          { return "FileEntry" }
func (fe *FileEntry) String() string        { return fe.AbsPath }

func (fe *FileEntry) Attr(name string) (strlk.Value, error) {
	switch name {
	case "abspath":
		return strlk.String(fe.AbsPath), nil
	case "basename":
		return strlk.String(fe.BaseName), nil
	case "dir":
		return strlk.String(fe.Dir), nil
	case "ext":
		return strlk.String(fe.Ext), nil
	case "name":
		return strlk.String(fe.Name), nil
	case "path":
		return strlk.String(fe.Path), nil
	}
	return starlark.None, fmt.Errorf("Unknown attr: %s", name)
}

func (fe *FileEntry) AttrNames() []string {
	return []string{
		"abspath",
		"basename",
		"dir",
		"ext",
		"name",
		"path",
	}
}

func (fe *FileEntry) Init(path string) (err error) {
	fe.Path = path
	fe.Name = filepath.Base(path)
	fe.Dir = filepath.Dir(path)
	fe.Ext = filepath.Ext(path)
	fe.BaseName = fe.Name[:len(fe.Name)-len(fe.Ext)]

	fe.AbsPath, err = filepath.Abs(path)
	if err != nil {
		fe.AbsPath = ""
		return err
	}

	return nil
}

func NewFileEntry(path string) (result *FileEntry, err error) {
	result = &FileEntry{}
	if err = result.Init(path); err != nil {
		return result, err
	}

	return result, nil
}

// parse
func parse(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		result *FileEntry
		path   string
	)

	if err = starlark.UnpackArgs("parse", args, kwargs, "path", &path); err != nil {
		return nil, err
	}

	result, _ = NewFileEntry(path)
	return result, nil
}
