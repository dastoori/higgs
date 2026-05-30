//go:build !windows

package higgs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// IsHidden checks whether "FileHide.Path" is hidden or not
func (fh *FileHide) IsHidden() (bool, error) {
	if fh.DarwinNative {
		return darwinHasHidden(fh.Path)
	}

	f, err := os.Stat(fh.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("\"%s\" does not exist", fh.Path)
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
			return fmt.Errorf("\"%s\" does not exist", fh.Path)
		}
		return fmt.Errorf("something went wrong getting file stat: \"%s\"", err)
	}

	// macOS native hiding toggles the UF_HIDDEN flag without renaming the file
	if fh.DarwinNative {
		return darwinSetHidden(fh.Path, hidden)
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

	// Guard against producing an empty or dot-only destination name
	if base := filepath.Base(dstName); base == "" || base == "." || base == ".." {
		return fmt.Errorf("cannot rename \"%s\" to an invalid name \"%s\"", fh.Path, dstName)
	}

	// Check destination file
	if !fh.UnixOverwrite {
		if _, statErr := os.Stat(dstName); statErr == nil {
			return fmt.Errorf("\"%s\" already exists\nSet the `UnixOverwrite` flag to skip this check", dstName)
		} else if !os.IsNotExist(statErr) {
			return fmt.Errorf("something went wrong checking destination \"%s\": \"%s\"", dstName, statErr)
		}
	}

	err = os.Rename(fh.Path, dstName)
	if err != nil {
		return fmt.Errorf("something went wrong renaming the \"%s\" to \"%s\": \"%s\"", fh.Path, dstName, err)
	}
	fh.Path = dstName

	return nil
}
