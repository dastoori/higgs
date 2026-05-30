package higgs_test

import (
	"fmt"

	"github.com/dastoori/higgs"
)

// ExampleHide hides a file using the package-level helper. On Unix the file is
// renamed with a leading dot and the returned path holds the new name; on
// Windows/macOS-native the path is unchanged.
func ExampleHide() {
	newPath, err := higgs.Hide("foo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("hidden:", newPath)
}

// ExampleUnhide unhides a file using the package-level helper.
func ExampleUnhide() {
	newPath, err := higgs.Unhide(".foo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("unhidden:", newPath)
}

// ExampleIsHidden reports whether a path is currently hidden.
func ExampleIsHidden() {
	hidden, err := higgs.IsHidden("foo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("hidden:", hidden)
}

// ExampleNewFileHide builds a FileHide with options and hides a file through
// the method API. UnixOverwriteOption lets the rename overwrite an existing
// destination file on Unix.
func ExampleNewFileHide() {
	fh := higgs.NewFileHide("bar.txt", higgs.UnixOverwriteOption(true))
	if err := fh.Hide(); err != nil {
		fmt.Println(err)
		return
	}

	// On Unix fh.Path now points at the renamed (dot-prefixed) file.
	fmt.Println("hidden:", fh.Path)
}

// ExampleDarwinNativeOption hides a file on macOS using the native UF_HIDDEN
// flag instead of renaming it, so fh.Path stays unchanged. The option returns
// an error on non-macOS unix systems.
func ExampleDarwinNativeOption() {
	fh := higgs.NewFileHide("baz.txt", higgs.DarwinNativeOption(true))
	if err := fh.Hide(); err != nil {
		fmt.Println(err)
		return
	}

	// fh.Path is unchanged because the file was not renamed.
	fmt.Println("hidden:", fh.Path)
}
