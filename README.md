# go-fshide

A tiny GO library to hide/show files and directories

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/dastoori/go-fshide/go)](https://github.com/dastoori/go-fshide/actions/workflows/go.yml)
[![Codecov](https://img.shields.io/codecov/c/github/dastoori/go-fshide)](https://app.codecov.io/gh/dastoori/go-fshide/)
[![GitHub release](https://img.shields.io/github/v/release/dastoori/go-fshide)](https://github.com/dastoori/go-fshide/releases)<br/>
[![Go Reference](https://pkg.go.dev/badge/github.com/dastoori/go-fshide.svg)](https://pkg.go.dev/github.com/dastoori/go-fshide)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dastoori/go-fshide)
[![GitHub](https://img.shields.io/github/license/dastoori/go-fshide)](https://github.com/dastoori/go-fshide/blob/master/LICENSE)

## Installation

```sh
$ go get github.com/dastoori/go-fshide
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/dastoori/go-fshide"
)

func main() {
	err := fshide.Hide("./path-to-hide", true)
	
	if err != nil {
		fmt.Println(err)
	}
}
```

## TODO

- [ ] Windows support
