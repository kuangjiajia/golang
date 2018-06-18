package main

import (
	"os"
	kjj "fmt"
)
// 忽略包 	_  pkgName
// 包别名  pkgName ""
// . pkgName

func main() {
	list := os.Args
	n := len(list)
	kjj.Println("n = ", n)
	kjj.Println("list = ", list)
	kjj.Println("list = ", list)
}