package fs

import (
	"fmt"
	"os"
	pathlib "path"
)

func AssertParentDir(path string, mkdir bool) error {
	parentDir := pathlib.Dir(path)
	if _, err := os.Stat(parentDir); os.IsNotExist(err) {
		if mkdir {
			if err = os.MkdirAll(parentDir, os.ModePerm); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("parent dir: \"%s\" for path doesn't exist", parentDir)
		}
	}

	return nil
}
