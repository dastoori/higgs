package higgs

// FileHidener implements the FileHide methods
type FileHidener interface {
	IsHidden() (bool, error)
	Hide() error
	Unhide() error
}

// FileHide object that holds higgs configs
type FileHide struct {
	Path          string
	UnixOverwrite bool
}

// FileHideOption type that holds a FileHide option
type FileHideOption func(*FileHide)

// NewFileHide makes new FileHide instance
func NewFileHide(path string, options ...FileHideOption) *FileHide {
	fh := &FileHide{
		Path:          path,
		UnixOverwrite: false,
	}

	for _, option := range options {
		option(fh)
	}

	return fh
}

// UnixOverwriteOption allows the renaming process to overwrite existing file (unix option)
func UnixOverwriteOption(value bool) FileHideOption {
	return func(fh *FileHide) {
		fh.UnixOverwrite = value
	}
}

// IsHidden checks whether "path" is hidden or not
func IsHidden(path string) (bool, error) {
	return NewFileHide(path).IsHidden()
}

// Hide makes file or directory hidden
func Hide(path string) (string, error) {
	fh := NewFileHide(path)
	err := fh.Hide()
	if err != nil {
		return "", err
	}
	return fh.Path, nil
}

// Unhide makes file or directory unhidden
func Unhide(path string) (string, error) {
	fh := NewFileHide(path)
	err := fh.Unhide()
	if err != nil {
		return "", err
	}
	return fh.Path, nil
}
