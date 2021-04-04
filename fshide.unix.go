// +build !windows

package fshide

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// IsHidden checks whether "FsHide.Path" is hidden or not
func (fh *FsHide) IsHidden() (bool, error) {
	f, err := os.Stat(fh.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("\"%s\" is not exists", fh.Path)
		}
		return false, fmt.Errorf("Something went wrong getting file stat: \"%s\"", err)
	}

	if strings.HasPrefix(f.Name(), ".") {
		return true, nil
	}

	return false, nil
}

// Hide makes file or directory hidden or unhidden
func (fh *FsHide) Hide(hidden bool) (err error) {
	srcFile, err := os.Stat(fh.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("\"%s\" is not exists", fh.Path)
		}
		return fmt.Errorf("Something went wrong getting file stat: \"%s\"", err)
	}

	if strings.HasPrefix(srcFile.Name(), ".") == hidden {
		// Nothing to do
		return nil
	}

	// Generate destination name
	var dstName string
	if hidden {
		dstName = filepath.Join(filepath.Dir(fh.Path), "."+filepath.Base(fh.Path))
	} else {
		dstName = filepath.Join(filepath.Dir(fh.Path), strings.TrimPrefix(filepath.Base(fh.Path), "."))
	}

	// Check destination file
	if !fh.Overwrite {
		_, err = os.Stat(dstName)
		if err == nil {
			return fmt.Errorf("\"%s\" already exists\nSet the `Overwrite` flag to skip this check", dstName)
		}
	}

	err = os.Rename(fh.Path, dstName)
	if err != nil {
		return fmt.Errorf("Something went wrong renaming the \"%s\" to \"%s\": \"%s\"", fh.Path, dstName, err)
	}

	return nil
}
