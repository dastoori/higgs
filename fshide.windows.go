// +build windows

package fshide

import (
	"fmt"
)

// IsHidden checks whether "FsHide.Path" is hidden or not
func (fh *FsHide) IsHidden() (bool, error) {
	// TODO implement windows support
	return false, fmt.Errorf("Not Implemented")
}

// Hide makes file or directory hidden or unhidden
func (fh *FsHide) Hide(hidden bool) error {
	// TODO implement windows support
	return fmt.Errorf("Not Implemented")
}
