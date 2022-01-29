//go:build windows
// +build windows

package higgs

import (
	"fmt"
	"os"
	"syscall"
)

func getFileAttrs(path string) (uint32, *uint16, error) {
	utf16PtrPath, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return 0, nil, fmt.Errorf("something went wrong getting path's UTF16 pointer: \"%s\"", err)
	}
	attrs, err := syscall.GetFileAttributes(utf16PtrPath)
	return attrs, utf16PtrPath, err
}

// IsHidden checks whether "FileHide.Path" is hidden or not
func (fh *FileHide) IsHidden() (bool, error) {
	attrs, _, err := getFileAttrs(fh.Path)
	if err != nil {
		return false, fmt.Errorf("something went wrong getting the file attributes: \"%s\"", err)
	}
	if attrs&syscall.FILE_ATTRIBUTE_HIDDEN > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// Hide makes file or directory hidden
func (fh *FileHide) Hide() error {
	return fh.hide(true)
}

// Unhide makes file or directory unhidden
func (fh *FileHide) Unhide() error {
	return fh.hide(false)
}

func (fh *FileHide) hide(hidden bool) error {
	_, err := os.Stat(fh.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("\"%s\" is not exists", fh.Path)
		}
		return fmt.Errorf("something went wrong getting file stat: \"%s\"", err)
	}

	attrs, utf16PtrPath, err := getFileAttrs(fh.Path)
	if err != nil {
		return fmt.Errorf("something went wrong getting the file attributes: \"%s\"", err)
	}

	var newAttrs uint32
	if hidden {
		if attrs&syscall.FILE_ATTRIBUTE_HIDDEN > 0 {
			return nil
		}
		// Add hidden attribute to file's current attributes
		newAttrs = attrs | syscall.FILE_ATTRIBUTE_HIDDEN
	} else {
		if attrs&syscall.FILE_ATTRIBUTE_HIDDEN == 0 {
			return nil
		}
		newAttrs = attrs - (attrs & syscall.FILE_ATTRIBUTE_HIDDEN)
	}
	err = syscall.SetFileAttributes(utf16PtrPath, newAttrs)
	if err != nil {
		return fmt.Errorf("something went wrong setting the file attributes: \"%s\"", err)
	}

	return nil
}
