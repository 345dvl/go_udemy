package main

import "fmt"

func main() {
	f:= func(x, y int) int {
		return x + y
	}

	fmt.Println(f(3, 4))

	i := func(x, y int) int {
		return x + y
	}(5, 4)
	fmt.Println(i)
}