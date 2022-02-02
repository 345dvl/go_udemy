package main

import (
	."fmt"
	f "fmt"
	"foo/foo"
)

// スコープ

func main() {
	Println(foo.Max)
	// fmt.Println(foo.min)

	f.Println(foo.ReturnMin())
}