package fs

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"go.starlark.net/starlark"
	strlk "go.starlark.net/starlark"
)

// {{{1 FileEntry
type FileEntryKind string

const (
	FileEntryDir  FileEntryKind = "Dir"
	FileEntryFile               = "File"
)

type FileEntry struct {
	Kind          FileEntryKind
	AbsPath       string
	BaseDir       string
	BaseName      string
	IsSymlink     bool
	SymlinkTarget string
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
	return map[string]strlk.Value{
		"kind":       strlk.String(fe.Kind),
		"path":       strlk.String(fe.AbsPath),
		"dir":        strlk.String(fe.BaseDir),
		"name":       strlk.String(fe.BaseName),
		"is_symlink": strlk.Bool(fe.IsSymlink),
	}[name], nil
}

func (fe *FileEntry) AttrNames() []string {
	return []string{"kind", "path", "dir", "name", "is_symlink"}
}

func FromPath(p string) (*FileEntry, error) {
	result := &FileEntry{}

	fi, err := os.Lstat(p)
	if err != nil {
		return nil, err
	}

	if fi.IsDir() {
		result.Kind = FileEntryDir
	} else {
		result.Kind = FileEntryFile
	}

	result.AbsPath, err = filepath.Abs(p)
	if err != nil {
		return nil, err
	}

	result.BaseName = filepath.Base(p)
	result.BaseDir = filepath.Dir(p)

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
		l:    []*FileEntry{initialPath},
	}
}

// {{{1 fileIterator
type fileIterator struct {
	path strlk.String
	l    []*FileEntry
	i    int
}

var (
	_ strlk.Iterator = &fileIterator{}
)

func (it *fileIterator) advance() {
	var e *FileEntry
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
