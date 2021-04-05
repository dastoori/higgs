// +build !windows

package higgs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// IsHidden checks whether "Hide.Path" is hidden or not
func (h *FileHide) IsHidden() (bool, error) {
	f, err := os.Stat(h.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("\"%s\" is not exists", h.Path)
		}
		return false, fmt.Errorf("something went wrong getting file stat: \"%s\"", err)
	}

	if strings.HasPrefix(f.Name(), ".") {
		return true, nil
	}

	return false, nil
}

// Hide makes file or directory hidden
func (h *FileHide) Hide() (err error) {
	return h.hide(true)
}

// Unhide makes file or directory unhidden
func (h *FileHide) Unhide() (err error) {
	return h.hide(false)
}

func (h *FileHide) hide(hidden bool) (err error) {
	srcFile, err := os.Stat(h.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("\"%s\" is not exists", h.Path)
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
		dstName = filepath.Join(filepath.Dir(h.Path), "."+filepath.Base(h.Path))
	} else {
		dstName = filepath.Join(filepath.Dir(h.Path), strings.TrimPrefix(filepath.Base(h.Path), "."))
	}

	// Check destination file
	if !h.Overwrite {
		_, err = os.Stat(dstName)
		if err == nil {
			return fmt.Errorf("\"%s\" already exists\nSet the `Overwrite` flag to skip this check", dstName)
		}
	}

	err = os.Rename(h.Path, dstName)
	if err != nil {
		return fmt.Errorf("something went wrong renaming the \"%s\" to \"%s\": \"%s\"", h.Path, dstName, err)
	}

	return nil
}
