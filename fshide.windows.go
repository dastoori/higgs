// +build windows

package fshide

import (
	"log"
)

func (fh *FsHide) IsHidden() (bool, error) {
	// TODO implement windows support
	return false, log.Error("Not Implemented")
}

func (fh *FsHide) Hide(hidden bool) error {
	// TODO implement windows support
	return log.Error("Not Implemented")
}
