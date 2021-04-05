# higgs

A tiny cross-platform Go library to hide/unhide files and directories

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/dastoori/higgs/go)](https://github.com/dastoori/higgs/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/dastoori/higgs/branch/master/graph/badge.svg?token=T1AJXSWI3F)](https://codecov.io/gh/dastoori/higgs)
[![Go Report Card](https://goreportcard.com/badge/github.com/dastoori/higgs)](https://goreportcard.com/report/github.com/dastoori/higgs)
[![GitHub release](https://img.shields.io/github/v/release/dastoori/higgs)](https://github.com/dastoori/higgs/releases)<br/>
[![Go Reference](https://pkg.go.dev/badge/github.com/dastoori/higgs.svg)](https://pkg.go.dev/github.com/dastoori/higgs)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dastoori/higgs)
![No Dependency](https://img.shields.io/badge/dependency-no-green)
[![GitHub](https://img.shields.io/github/license/dastoori/higgs)](https://github.com/dastoori/higgs/blob/master/LICENSE)

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
	err := higgs.Hide("./path-to-hide")
	
	if err != nil {
		fmt.Println(err)
	}
}
```

## TODO

- [ ] Windows support

## License

MIT