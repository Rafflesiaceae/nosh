package fs

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"go.starlark.net/starlark"
	strlk "go.starlark.net/starlark"
)

// {{{1 StatFileEntry
type StatFileEntryKind string

const (
	FileEntryDir  StatFileEntryKind = "Dir"
	FileEntryFile                   = "File"
)

type StatFileEntry struct {
	FileEntry
	Kind StatFileEntryKind
	// AbsPath       string
	// BaseDir       string
	// BaseName      string
	IsSymlink     bool
	SymlinkTarget string
}

var (
	_ strlk.Value    = &StatFileEntry{}
	_ strlk.HasAttrs = &StatFileEntry{}
)

func (fe *StatFileEntry) Freeze()               { /* @TODO */ }
func (fe *StatFileEntry) Hash() (uint32, error) { /* @TODO */ return 0, nil }
func (fe *StatFileEntry) Truth() strlk.Bool     { return true }
func (fe *StatFileEntry) Type() string          { return "StatFileEntry" }
func (fe *StatFileEntry) String() string        { return fe.AbsPath }

func (fe *StatFileEntry) Attr(name string) (result strlk.Value, err error) {
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

	case "kind":
		return strlk.String(fe.Kind), nil
	case "is_symlink":
		return strlk.Bool(fe.IsSymlink), nil
	}
	return starlark.None, fmt.Errorf("Unknown attr: %s", name)
}

func (fe *StatFileEntry) AttrNames() []string {
	return []string{
		"abspath",
		"basename",
		"dir",
		"ext",
		"name",
		"path",

		"kind",
		"is_symlink",
	}
}

func FromPath(p string) (*StatFileEntry, error) {
	result := &StatFileEntry{}

	fi, err := os.Lstat(p)
	if err != nil {
		return nil, err
	}

	if fi.IsDir() {
		result.Kind = FileEntryDir
	} else {
		result.Kind = FileEntryFile
	}

	_ = result.FileEntry.Init(p)

	if fi.Mode()&fs.ModeSymlink != 0 {
		result.IsSymlink = true
	}

	return result, nil
}

// {{{1 fileIterable
type fileIterable struct{ path strlk.String }

var (
	_ strlk.Iterable = &fileIterable{}
	_ strlk.Value    = &fileIterable{}
)

func (it *fileIterable) Freeze()               { /* @TODO */ }
func (it *fileIterable) Hash() (uint32, error) { /* @TODO */ return 0, nil }
func (it *fileIterable) Truth() strlk.Bool     { return true }
func (it *fileIterable) Type() string          { return "file-iterable" }
func (it *fileIterable) String() string        { return fmt.Sprintf("file-iterable(%s)", it.path) }

func (it *fileIterable) Iterate() strlk.Iterator {
	initialPath, err := FromPath(it.path.GoString())
	if err != nil {
		panic(err)
	}

	return &fileIterator{
		path: it.path,
		l:    []*StatFileEntry{initialPath},
	}
}

// {{{1 fileIterator
type fileIterator struct {
	path strlk.String
	l    []*StatFileEntry
	i    int
}

var (
	_ strlk.Iterator = &fileIterator{}
)

func (it *fileIterator) advance() {
	var e *StatFileEntry
	e = it.l[it.i]

	// info, err := os.Lstat(p.GoString())
	// e, err := FromPath(e)
	// if err != nil {
	// 	return
	// 	// err = fn(root, nil, err)
	// }

	it.l = it.l[1:] // remove first element

	switch e.Kind {
	case FileEntryDir:
		// @TODO @BUG also output further dirs, not only first dir and then only
		// files from then on

		files, err := os.ReadDir(e.AbsPath)
		if err != nil {
			return
		}

		for _, f := range files {
			ne, err := FromPath(filepath.Join(e.AbsPath, f.Name()))
			if err != nil {
				continue
			}
			it.l = append(it.l, ne)
		}
	case FileEntryFile:
	}
	// if e.Type == FileEntryDir {
	// 	err = walkDir(root, &statDirEntry{info}, fn)
	// }
	// if err == SkipDir {
	// 	return nil
	// }

	return
}

func (it *fileIterator) Next(p *strlk.Value) bool {
	// if it.i < it.l.Len() {
	if len(it.l) > 0 {
		*p = it.l[it.i]
		// *p = it.l.elems[it.i]
		// it.i++
		it.advance()
		return true
	}
	return false
}

func (it *fileIterator) Done() {
	// @TODO
	// if !it.l.frozen {
	// 	it.l.itercount--
	// }
}

// {{{1 fs.find

// @TODO use ↓ for globs for file-walking
// https://github.com/mattn/go-zglob.git

func visit(path string, di fs.DirEntry, err error) error {
	fmt.Printf("Visited: %s\n", path)

	return nil
}

func find(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		path          = strlk.String("")
		followSymlink = strlk.Bool(true)
	)
	if err = starlark.UnpackArgs("find", args, kwargs, "path", &path, "followSymlinks?", &followSymlink); err != nil {
		return nil, err
	}

	return &fileIterable{path: path}, nil
}
