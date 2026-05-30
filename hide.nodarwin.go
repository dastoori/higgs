//go:build !windows && !darwin

package higgs

import "fmt"

// The native UF_HIDDEN file flag only exists on macOS. On other unix systems the
// DarwinNative option is unsupported, so these stubs report a clear error. They
// exist so hide.unix.go (built for every non-windows OS) compiles everywhere.

func darwinHasHidden(path string) (bool, error) {
	return false, fmt.Errorf("native (UF_HIDDEN) hiding is only supported on macOS")
}

func darwinSetHidden(path string, hidden bool) error {
	return fmt.Errorf("native (UF_HIDDEN) hiding is only supported on macOS")
}
