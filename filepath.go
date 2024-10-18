package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/wordl.go"))
	fmt.Println(filepath.Base("hello/wordl.go"))
	fmt.Println(filepath.Ext("hello/wordl.go"))
	fmt.Println(filepath.IsAbs("hello/wordl.go"))
	fmt.Println(filepath.IsLocal("hello/wordl.go"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))

}
