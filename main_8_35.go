package main

import "fmt"

func main() {
	var a int = 2
	if a == 2 {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}

	if b := 2; b == 2 {
		fmt.Println(b)
	}
}