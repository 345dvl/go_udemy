package main

import "fmt"

func anything(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println((v + 10000))
	}
}

func main() {
	// 型アサーション
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i := x.(int)
	fmt.Println(i + 2)

	// f, isFloat64 := x.(float64)
	// fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("none")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, isInt)
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("i don't know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("i dont know")
	}
}