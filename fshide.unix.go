// +build !windows

package fshide

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// IsHidden checks whether "fh.Path" is hidden or not
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

// Hide makes file or directory hidden or not hidden
func (fsh *FsHide) Hide(hidden bool) (err error) {
	srcFile, err := os.Stat(fsh.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("\"%s\" is not exists", fsh.Path)
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
		dstName = filepath.Join(filepath.Dir(fsh.Path), "."+filepath.Base(fsh.Path))
	} else {
		dstName = filepath.Join(filepath.Dir(fsh.Path), strings.TrimPrefix(filepath.Base(fsh.Path), "."))
	}

	// Check destination file
	if !fsh.Overwrite {
		_, err = os.Stat(dstName)
		if err == nil {
			return fmt.Errorf("\"%s\" already exists\nSet the `Overwrite` flag to skip this check", dstName)
		}
	}

	err = os.Rename(fsh.Path, dstName)
	if err != nil {
		return fmt.Errorf("Something went wrong renaming the \"%s\"", fsh.Path)
	}

	return nil
}
