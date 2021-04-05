/*
A tiny cross-platform Go library to hide/unhide files and directories

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
*/
package higgs
