# higgs

A tiny cross-platform Go library to hide or unhide files and directories.

**Supported OSs:** All unix based OSs (Tested on Ubuntu, MacOS), Windows

[![unix](https://github.com/dastoori/higgs/actions/workflows/unix.yml/badge.svg)](https://github.com/dastoori/higgs/actions/workflows/unix.yml)
[![windows](https://github.com/dastoori/higgs/actions/workflows/windows.yml/badge.svg)](https://github.com/dastoori/higgs/actions/workflows/windows.yml)
[![codecov](https://codecov.io/gh/dastoori/higgs/branch/master/graph/badge.svg?token=T1AJXSWI3F)](https://codecov.io/gh/dastoori/higgs)
[![GitHub release](https://img.shields.io/github/v/release/dastoori/higgs?sort=semver)](https://github.com/dastoori/higgs/releases)<br/>
[![Go Reference](https://pkg.go.dev/badge/github.com/dastoori/higgs.svg)](https://pkg.go.dev/github.com/dastoori/higgs)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dastoori/higgs)
[![Go Report Card](https://goreportcard.com/badge/github.com/dastoori/higgs)](https://goreportcard.com/report/github.com/dastoori/higgs)
[![GitHub](https://img.shields.io/github/license/dastoori/higgs)](https://github.com/dastoori/higgs/blob/master/LICENSE)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  

## Installation

```sh
$ go get github.com/dastoori/higgs
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/dastoori/higgs"
)

func main() {
	path := "foo.txt"

	// Hiding a file
	newPath, err := higgs.Hide(path)
	// NOTE: On Unix after hiding the file the file name
	// will be changed to `.foo.txt` and `newPath` contains
	// the new file name
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unhiding a file
	newPath, err = higgs.Unhide(newPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("unhidden:", newPath)

	// Setting the unix overwrite option (disabled by default)
	fh := higgs.NewFileHide(".bar.txt", higgs.UnixOverwriteOption(true))

	// NOTE: On Unix if a `bar.txt` file exists, it will be
	// overwritten after unhiding `.bar.txt`
	if err = fh.Unhide(); err != nil {
		fmt.Println(err)
		return
	}

	// NOTE: `fh.Path` contains the new file name
	fmt.Println("unhidden:", fh.Path)
}
```

### Checking whether a path is hidden

```go
hidden, err := higgs.IsHidden("foo.txt")
if err != nil {
	fmt.Println(err)
	return
}

fmt.Println("hidden:", hidden)
```

## How hiding works per platform

higgs hides files differently depending on the operating system:

- **Unix (Linux, \*BSD, macOS):** a file is hidden by renaming it with a leading dot
  (`foo.txt` → `.foo.txt`). Because the file is renamed, `FileHide.Path` (and the
  `string` returned by the package-level `Hide`/`Unhide` helpers) is updated to the new
  name. Use `UnixOverwriteOption(true)` to allow the rename to overwrite an existing file.
- **Windows:** a file is hidden by toggling the `FILE_ATTRIBUTE_HIDDEN` attribute. The
  file is **not** renamed, so `FileHide.Path` is unchanged.
- **macOS (native, opt-in):** with `DarwinNativeOption(true)`, higgs sets/clears the
  `UF_HIDDEN` file flag (via `chflags`) instead of renaming. Finder hides the file, but
  it is **not** renamed, so `FileHide.Path` is unchanged (like Windows). This option is
  only supported on macOS — it returns an error on other Unix systems and has no effect
  on Windows.

`Hide`, `Unhide` and the `FileHide` methods are idempotent: hiding a file that is already
hidden (or unhiding one that is already visible) is a no-op and returns no error.

## License

MIT