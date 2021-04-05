// +build windows

package higgs

import (
	"fmt"
)

// IsHidden checks whether "FileHide.Path" is hidden or not
func (h *FileHide) IsHidden() (bool, error) {
	// TODO implement windows support
	return false, fmt.Errorf("not Implemented")
}

// Hide makes file or directory hidden or unhidden
func (h *FileHide) Hide() error {
	// TODO implement windows support
	return fmt.Errorf("not Implemented")
}

// Unhide makes file or directory hidden or unhidden
func (h *FileHide) Unhide() error {
	// TODO implement windows support
	return fmt.Errorf("not Implemented")
}
