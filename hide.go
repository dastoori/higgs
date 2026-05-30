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
	DarwinNative  bool
}

// FileHideOption type that holds a FileHide option
type FileHideOption func(*FileHide)

// NewFileHide makes new FileHide instance
func NewFileHide(path string, options ...FileHideOption) *FileHide {
	fh := &FileHide{
		Path:          path,
		UnixOverwrite: false,
		DarwinNative:  false,
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

// DarwinNativeOption switches macOS to hide/unhide using the native UF_HIDDEN
// file flag instead of renaming with a leading dot (macOS option).
//
// When enabled, Hide/Unhide set or clear the UF_HIDDEN flag via chflags, so the
// file is hidden from Finder without being renamed and "FileHide.Path" stays
// unchanged (similar to the Windows behavior). It has no effect on Windows and
// returns an error on non-macOS unix systems.
func DarwinNativeOption(value bool) FileHideOption {
	return func(fh *FileHide) {
		fh.DarwinNative = value
	}
}

// IsHidden checks whether "path" is hidden or not
func IsHidden(path string) (bool, error) {
	return NewFileHide(path).IsHidden()
}

// Hide makes a file or directory hidden and returns its path. On Unix the file
// is renamed with a leading dot, so the returned path differs from the input;
// on Windows the path is unchanged. Hiding an already-hidden path is a no-op.
func Hide(path string) (string, error) {
	fh := NewFileHide(path)
	err := fh.Hide()
	if err != nil {
		return "", err
	}
	return fh.Path, nil
}

// Unhide makes a file or directory unhidden and returns its path. On Unix the
// leading dot is removed, so the returned path differs from the input; on
// Windows the path is unchanged. Unhiding an already-visible path is a no-op.
func Unhide(path string) (string, error) {
	fh := NewFileHide(path)
	err := fh.Unhide()
	if err != nil {
		return "", err
	}
	return fh.Path, nil
}
