package main

import "fmt"

func Divx(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func NoReturn() {
	fmt.Println("No Return")
	return
}

func main(){
	q1, r1 := Divx(9, 2)
	fmt.Println(q1, r1)

	fmt.Println(Double(100))

	NoReturn()
}