//go:build darwin

package higgs

import (
	"path/filepath"
	"testing"
)

// These tests exercise the native UF_HIDDEN code path (DarwinNativeOption) and
// only run on macOS, where the GitHub Actions runner provides a real darwin
// filesystem. They reuse the shared TestMain/tmpDir setup from hide.unix_test.go.

func TestDarwinNativeIsHiddenWhenNotHidden(t *testing.T) {
	path := filepath.Join(tmpDir, "a")
	hidden, err := NewFileHide(path, DarwinNativeOption(true)).IsHidden()

	if err != nil {
		t.Errorf("error: \"%s\"", err)
	}

	if hidden {
		t.Errorf("file should not be hidden but the output says otherwise")
	}
}

func TestDarwinNativeIsHiddenWhenNotExists(t *testing.T) {
	path := filepath.Join(tmpDir, "notexists")
	_, err := NewFileHide(path, DarwinNativeOption(true)).IsHidden()

	if err == nil {
		t.Errorf("expected an error for a non-existent path")
	}
}

func TestDarwinNativeHidesAndUnhidesWithoutRenaming(t *testing.T) {
	path := filepath.Join(tmpDir, "native")
	fh := NewFileHide(path, DarwinNativeOption(true))

	if err := fh.Hide(); err != nil {
		t.Fatalf("error hiding: \"%s\"", err)
	}

	if fh.Path != path {
		t.Errorf("native hiding must not rename the file: \"%s\"", fh.Path)
	}

	hidden, err := fh.IsHidden()
	if err != nil {
		t.Errorf("error: \"%s\"", err)
	}
	if !hidden {
		t.Errorf("file should be hidden but it is not: \"%s\"", path)
	}

	// Hiding again is a no-op
	if err := fh.Hide(); err != nil {
		t.Errorf("re-hiding should be a no-op: \"%s\"", err)
	}

	if err := fh.Unhide(); err != nil {
		t.Fatalf("error unhiding: \"%s\"", err)
	}

	if fh.Path != path {
		t.Errorf("native unhiding must not rename the file: \"%s\"", fh.Path)
	}

	hidden, err = fh.IsHidden()
	if err != nil {
		t.Errorf("error: \"%s\"", err)
	}
	if hidden {
		t.Errorf("file should not be hidden but it is: \"%s\"", path)
	}

	// Unhiding again is a no-op
	if err := fh.Unhide(); err != nil {
		t.Errorf("re-unhiding should be a no-op: \"%s\"", err)
	}
}

func TestDarwinNativeHideWhenNotExists(t *testing.T) {
	path := filepath.Join(tmpDir, "notexists")
	err := NewFileHide(path, DarwinNativeOption(true)).Hide()

	if err == nil {
		t.Errorf("expected an error for a non-existent path")
	}
}
