//go:build !windows
// +build !windows

package higgs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// IsHidden checks whether "FileHide.Path" is hidden or not
func (fh *FileHide) IsHidden() (bool, error) {
	f, err := os.Stat(fh.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("\"%s\" is not exists", fh.Path)
		}
		return false, fmt.Errorf("something went wrong getting file stat: \"%s\"", err)
	}

	if strings.HasPrefix(f.Name(), ".") {
		return true, nil
	}

	return false, nil
}

// Hide makes file or directory hidden
func (fh *FileHide) Hide() error {
	return fh.hide(true)
}

// Unhide makes file or directory unhidden
func (fh *FileHide) Unhide() error {
	return fh.hide(false)
}

func (fh *FileHide) hide(hidden bool) (err error) {
	srcFile, err := os.Stat(fh.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("\"%s\" is not exists", fh.Path)
		}
		return fmt.Errorf("something went wrong getting file stat: \"%s\"", err)
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
	if !fh.UnixOverwrite {
		_, err = os.Stat(dstName)
		if err == nil {
			return fmt.Errorf("\"%s\" already exists\nSet the `UnixOverwrite` flag to skip this check", dstName)
		}
	}

	err = os.Rename(fh.Path, dstName)
	if err != nil {
		return fmt.Errorf("something went wrong renaming the \"%s\" to \"%s\": \"%s\"", fh.Path, dstName, err)
	}
	fh.Path = dstName

	return nil
}
