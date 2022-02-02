package main75

import "fmt"

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main75() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))
}