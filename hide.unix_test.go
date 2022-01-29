//go:build !windows
// +build !windows

package higgs

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var tmpDir string

func touch(path, content string) {
	path = filepath.FromSlash(path)
	dir := filepath.Dir(path)
	if dir != "" {
		os.MkdirAll(filepath.Join(tmpDir, dir), 0755)
	}
	ioutil.WriteFile(filepath.Join(tmpDir, path), []byte(content), 0644)
}

func TestMain(t *testing.M) {
	tmpDir, _ = ioutil.TempDir("", "higgs*")
	touch("a", "a")
	touch(".b", "b")
	touch("c/c.a", "c.a")
	touch("d", "d")
	touch(".d", "d hidden")

	code := t.Run()
	defer func() {
		os.RemoveAll(tmpDir)
		os.Exit(code)
	}()
}

func TestIsHiddenWhenNotHidden(t *testing.T) {
	hidden, err := IsHidden(filepath.Join(tmpDir, "a"))

	if err != nil {
		t.Errorf(`Error: "%s"`, err)
	}
	if hidden == true {
		t.Errorf("wrong output, file is not hidden but the output says otherwise")
	}
}

func TestIsHiddenWhenHidden(t *testing.T) {
	hidden, err := IsHidden(filepath.Join(tmpDir, ".b"))

	if err != nil {
		t.Errorf(`Error: "%s"`, err)
	}
	if hidden == false {
		t.Errorf("wrong output, file is hidden but the output says otherwise")
	}
}

func TestIsHiddenWhenNotExists(t *testing.T) {
	hidden, err := IsHidden(filepath.Join(tmpDir, "notexists"))

	if err == nil {
		t.Errorf("no error")
	}
	if hidden == true {
		t.Errorf("wrong output")
	}
}

func TestHiddenHidesWhenAlreadyHidden(t *testing.T) {
	path := filepath.Join(tmpDir, ".b")
	newPath, err := Hide(path)

	if err != nil {
		t.Errorf("error: \"%s\"", err)
	}

	if path != newPath {
		t.Errorf("the file wrongly renamed: \"%s\"", newPath)
	}

	_, err = os.Stat(path)

	if err != nil {
		t.Errorf("error: \"%s\"", err)
	}
}

func TestUnhideUnhidesWhenAlreadyNotHidden(t *testing.T) {
	path := filepath.Join(tmpDir, "a")
	newPath, err := Unhide(path)

	if err != nil {
		t.Errorf("error: \"%s\"", err)
	}

	if path != newPath {
		t.Errorf("the file wrongly renamed: \"%s\"", newPath)
	}

	_, err = os.Stat(path)

	if err != nil {
		t.Errorf("error: \"%s\"", err)
	}
}

func TestHideWhenNotExists(t *testing.T) {
	path := filepath.Join(tmpDir, "notexists")
	newPath, err := Hide(path)

	if err == nil {
		t.Errorf("error: \"%s\"", err)
	}

	if path == "" {
		t.Errorf("the file wrongly renamed: \"%s\"", newPath)
	}
}

func TestHiddenHidesFile(t *testing.T) {
	path := filepath.Join(tmpDir, "a")
	newPath, err := Hide(path)

	if err != nil {
		t.Errorf("error: \"%s\"", err)
	}

	if path == newPath || !strings.HasSuffix(newPath, ".a") {
		t.Errorf("the file wrongly renamed: \"%s\"", newPath)
	}

	_, err = os.Stat(path)

	if err != nil && errors.Is(err, os.ErrExist) {
		t.Errorf("the file is still exists: \"%s\"", err)
	}
}

func TestUnhideUnhidesFile(t *testing.T) {
	path := filepath.Join(tmpDir, ".b")
	newPath, err := Unhide(path)

	if err != nil {
		t.Errorf("error: \"%s\"", err)
	}

	if path == newPath || !strings.HasSuffix(newPath, "b") {
		t.Errorf("the file wrongly renamed: \"%s\"", newPath)
	}

	_, err = os.Stat(path)

	if err != nil && errors.Is(err, os.ErrExist) {
		t.Errorf("the file is still exists: \"%s\"", err)
	}
}

func TestHiddenHidesDirectory(t *testing.T) {
	path := filepath.Join(tmpDir, "c")
	newPath, err := Hide(path)

	if err != nil {
		t.Errorf("error: \"%s\"", err)
	}

	if path == newPath || !strings.HasSuffix(newPath, ".c") {
		t.Errorf("the file wrongly renamed: \"%s\"", newPath)
	}

	_, err = os.Stat(path)

	if err != nil && errors.Is(err, os.ErrExist) {
		t.Errorf("the directory is still exists: \"%s\"", err)
	}
}

func TestHiddenCantHidesNoOverwrite(t *testing.T) {
	path := filepath.Join(tmpDir, "d")
	fh := NewFileHide(path)
	err := fh.Hide()

	if err == nil {
		t.Errorf("error: \"%s\"", err)
	}

	if path != fh.Path {
		t.Errorf("the file wrongly renamed: \"%s\"", fh.Path)
	}

	_, err = os.Stat(path)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		t.Errorf("the file is not exists: \"%s\"", err)
	}
}

func TestHiddenHidesWithOverwrite(t *testing.T) {
	path := filepath.Join(tmpDir, "d")
	fh := NewFileHide(path, UnixOverwriteOption(true))
	err := fh.Hide()

	if err != nil {
		t.Errorf("error: \"%s\"", err)
	}

	if path == fh.Path || !strings.HasSuffix(fh.Path, ".d") {
		t.Errorf("the file wrongly renamed: \"%s\"", fh.Path)
	}

	_, err = os.Stat(path)

	if err == nil {
		t.Errorf("the file is still exists: \"%s\"", err)
	}
}
