# FsHide

A tiny GO library to hide/show files and directories

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
