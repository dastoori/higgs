// +build windows

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
	touch("b", "b")
	touch("c/c.a", "c.a")
	touch("d", "d")

	code := t.Run()
	defer func() {
		os.RemoveAll(tmpDir)
		os.Exit(code)
	}()
}

func TestIsHidden(t *testing.T) {
	// TODO implement windows support
}

func TestHide(t *testing.T) {
	// TODO implement windows support
}
