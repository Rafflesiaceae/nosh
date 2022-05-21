package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AssertParentDir(path string, mkdir bool) error {
	parentDir := filepath.Dir(path)
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

func IsPathWithin(smallerPath string, biggerPath string) (bool, error) {
	biggerAbsPath, err := filepath.Abs(biggerPath)
	if err != nil {
		return false, err
	}

	smallerAbsPath, err := filepath.Abs(smallerPath)
	if err != nil {
		return false, err
	}

	if biggerAbsPath == smallerAbsPath {
		return true, nil
	}

	biggerPathParts := strings.Split(biggerAbsPath, string(os.PathSeparator))
	smallerPathParts := strings.Split(smallerAbsPath, string(os.PathSeparator))

	lenBiggerPathParts := len(biggerPathParts)
	for i, v := range smallerPathParts {
		if i >= lenBiggerPathParts || v != biggerPathParts[i] {
			return false, nil
		}
	}

	return true, nil
}
