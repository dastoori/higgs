// +build !windows

package fshide

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
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
	tmpDir, _ = ioutil.TempDir("", "fshide*")
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
		t.Errorf("Wrong output, file is not hidden but the output says otherwise")
	}
}

func TestIsHiddenWhenHidden(t *testing.T) {
	hidden, err := IsHidden(filepath.Join(tmpDir, ".b"))

	if err != nil {
		t.Errorf(`Error: "%s"`, err)
	}
	if hidden == false {
		t.Errorf("Wrong output, file is hidden but the output says otherwise")
	}
}

func TestIsHiddenWhenNotExists(t *testing.T) {
	hidden, err := IsHidden(filepath.Join(tmpDir, "notexists"))

	if err == nil {
		t.Errorf("No error")
	}
	if hidden == true {
		t.Errorf("Wrong output")
	}
}

func TestHideHidesWhenAlreadyHidden(t *testing.T) {
	path := filepath.Join(tmpDir, ".b")
	err := Hide(path, true)

	if err != nil {
		t.Errorf("Error: \"%s\"", err)
	}

	_, err = os.Stat(path)

	if err != nil {
		t.Errorf("Error: \"%s\"", err)
	}
}

func TestHideNotHidesWhenAlreadyNotHidden(t *testing.T) {
	path := filepath.Join(tmpDir, "a")
	err := Hide(path, false)

	if err != nil {
		t.Errorf("Error: \"%s\"", err)
	}

	_, err = os.Stat(path)

	if err != nil {
		t.Errorf("Error: \"%s\"", err)
	}
}

func TestHideWhenNotExists(t *testing.T) {
	err := Hide(filepath.Join(tmpDir, "notexists"), true)

	if err == nil {
		t.Errorf("Error: \"%s\"", err)
	}
}

func TestHideHidesFile(t *testing.T) {
	path := filepath.Join(tmpDir, "a")
	err := Hide(path, true)

	if err != nil {
		t.Errorf("Error: \"%s\"", err)
	}

	_, err = os.Stat(path)

	if err != nil && errors.Is(err, os.ErrExist) {
		t.Errorf("The file is still exists: \"%s\"", err)
	}
}

func TestHideUnhidesFile(t *testing.T) {
	path := filepath.Join(tmpDir, ".b")
	err := Hide(path, false)

	if err != nil {
		t.Errorf("Error: \"%s\"", err)
	}

	_, err = os.Stat(path)

	if err != nil && errors.Is(err, os.ErrExist) {
		t.Errorf("The file is still exists: \"%s\"", err)
	}
}

func TestHideHidesDirectory(t *testing.T) {
	path := filepath.Join(tmpDir, "c")
	err := Hide(path, true)

	if err != nil {
		t.Errorf("Error: \"%s\"", err)
	}

	_, err = os.Stat(path)

	if err != nil && errors.Is(err, os.ErrExist) {
		t.Errorf("The directory is still exists: \"%s\"", err)
	}
}

func TestHideCantHidesNoOverwrite(t *testing.T) {
	path := filepath.Join(tmpDir, "d")
	err := NewFsHide(path, false).Hide(true)

	if err == nil {
		t.Errorf("Error: \"%s\"", err)
	}

	_, err = os.Stat(path)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		t.Errorf("The file is not exists: \"%s\"", err)
	}
}

func TestHideHidesWithOverwrite(t *testing.T) {
	path := filepath.Join(tmpDir, "d")
	err := NewFsHide(path, true).Hide(true)

	if err == nil {
		t.Errorf("Error: \"%s\"", err)
	}

	_, err = os.Stat(path)

	if err != nil && errors.Is(err, os.ErrExist) {
		t.Errorf("The file is still exists: \"%s\"", err)
	}
}
