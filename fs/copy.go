package fs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"go.starlark.net/starlark"
)

func copyImpl(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		from  string
		to    string
		force bool = false
	)

	if err = starlark.UnpackArgs("copy", args, kwargs, "from", &from, "to", &to, "force?", &force); err != nil {
		return nil, err
	}

	if err = CopyPath(from, to, force); err != nil {
		return nil, err
	}

	return starlark.None, nil
}

func CopyPath(from string, to string, force bool) error {
	if _, err := os.Stat(from); os.IsNotExist(err) {
		return fmt.Errorf("from path: \"%s\" doesn't exist", from)
	}

	if !force {
		if _, err := os.Stat(to); !os.IsNotExist(err) {
			return fmt.Errorf("to path: \"%s\" already exists", to)
		}
	}

	if fi, err := os.Stat(from); err != nil {
		return err

	} else if fi.IsDir() {
		if err = CopyDir(from, to); err != nil {
			return err
		}
	} else {
		if err = CopyFile(from, to); err != nil {
			return err
		}
	}

	return nil
}

func CopyDir(from string, to string) error {
	var err error

	si, err := os.Stat(from)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(to, si.Mode()); err != nil {
		return err
	}

	dir, err := os.Open(from)
	if err != nil {
		return err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	for _, file := range files {
		childFrom := filepath.Join(from+string(os.PathSeparator), file.Name())
		childTo := filepath.Join(to+string(os.PathSeparator), file.Name())

		if file.IsDir() {
			if err = CopyDir(childFrom, childTo); err != nil {
				return err
			}
		} else {
			if err = CopyFile(childFrom, childTo); err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFile(from string, to string) error {
	var err error

	ln, err := os.Readlink(from)
	if err == nil {
		return os.Symlink(ln, to)
	}

	fromStream, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fromStream.Close()

	si, err := os.Stat(from)
	if err != nil {
		return err
	}

	// Create File
	toStream, err := os.OpenFile(to, os.O_RDWR|os.O_CREATE|os.O_TRUNC, si.Mode())
	if err != nil {
		return err
	}
	defer toStream.Close()

	_, err = io.Copy(toStream, fromStream)
	if err != nil {
		return err
	}

	return err
}
