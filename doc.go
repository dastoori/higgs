// Package higgs is a tiny cross-platform Go library to hide or unhide files and directories
//
// 	package main
//
// 	import (
// 		"fmt"
// 		"github.com/dastoori/higgs"
// 	)
//
// 	func main() {
// 		err := higgs.Hide("foo.txt")
//
// 		if err != nil {
// 			fmt.Println(err)
// 		}
//
// 		err = higgs.Unhide("foo.txt")
//
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}
package higgs
