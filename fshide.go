package fshide

type FsHidener interface {
	IsHidden() (bool, error)
	Hide(bool) error
}

type FsHide struct {
	Path      string
	Overwrite bool
}

// NewFsHide makes new FsHide instance
func NewFsHide(path string, overwrite bool) *FsHide {
	return &FsHide{Path: path}
}

// IsHidden checks whether "path" is hidden or not
func IsHidden(path string) (bool, error) {
	return NewFsHide(path, false).IsHidden()
}

// Hide makes file or directory hidden or unhidden
func Hide(path string, hidden bool) error {
	return NewFsHide(path, false).Hide(hidden)
}
