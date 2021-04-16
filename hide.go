package higgs

// FileHidener implements the FileHide methods
type FileHidener interface {
	IsHidden() (bool, error)
	Hide() error
	Unhide() error
}

// FileHide object that holds higgs configs
type FileHide struct {
	Path      string
	Overwrite bool
}

// NewFileHide makes new Hide instance
func NewFileHide(path string, overwrite bool) *FileHide {
	return &FileHide{
		Path:      path,
		Overwrite: overwrite,
	}
}

// IsHidden checks whether "path" is hidden or not
func IsHidden(path string) (bool, error) {
	return NewFileHide(path, false).IsHidden()
}

// Hide makes file or directory hidden
func Hide(path string) error {
	return NewFileHide(path, false).Hide()
}

// Unhide makes file or directory unhidden
func Unhide(path string) error {
	return NewFileHide(path, false).Unhide()
}
