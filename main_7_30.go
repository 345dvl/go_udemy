package main

import "fmt"

func CallFunction(f func()) {
	f()
}

func main() {
	CallFunction(func() {
		fmt.Println("I am a function")
	})
}