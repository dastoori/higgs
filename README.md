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
	}

	// Unhiding a file
	newPath, err = higgs.Unhide(newPath)
	
	if err != nil {
		fmt.Println(err)
	}

	// Setting unix overwrite option (disable by default)
	fh := NewFileHide(".bar.txt", UnixOverwriteOption(true))

	// NOTE: On Unix if a `bar.txt` file exists, it will be
	// overwritten after unhiding `.bar.txt`
	err := fh.Unhide()
	
	if err != nil {
		fmt.Println(err)
	}

	// NOTE: `fh.Path` contains the new file name
}
```

## License

MIT