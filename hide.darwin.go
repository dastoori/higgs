//go:build darwin

package higgs

import (
	"fmt"
	"os"
	"syscall"
)

// ufHidden is the macOS/BSD UF_HIDDEN file flag. Finder hides files that have
// it set. The standard "syscall" package does not export this constant on
// darwin, so it is defined here to keep higgs dependency free.
const ufHidden = 0x8000

// darwinHasHidden reports whether "path" has the UF_HIDDEN flag set.
func darwinHasHidden(path string) (bool, error) {
	var st syscall.Stat_t
	if err := syscall.Stat(path, &st); err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("\"%s\" does not exist", path)
		}
		return false, fmt.Errorf("something went wrong getting file stat: \"%s\"", err)
	}
	return st.Flags&ufHidden != 0, nil
}

// darwinSetHidden sets or clears the UF_HIDDEN flag on "path". It is idempotent:
// if the flag is already in the requested state it does nothing.
func darwinSetHidden(path string, hidden bool) error {
	var st syscall.Stat_t
	if err := syscall.Stat(path, &st); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("\"%s\" does not exist", path)
		}
		return fmt.Errorf("something went wrong getting file stat: \"%s\"", err)
	}

	flags := int(st.Flags)
	if hidden {
		if flags&ufHidden != 0 {
			return nil
		}
		flags |= ufHidden
	} else {
		if flags&ufHidden == 0 {
			return nil
		}
		flags &^= ufHidden
	}

	if err := syscall.Chflags(path, flags); err != nil {
		return fmt.Errorf("something went wrong setting the file flags: \"%s\"", err)
	}

	return nil
}
