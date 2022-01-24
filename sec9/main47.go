package main

import "fmt"

// スライス
// append make len cap

// func main() {
// 	// sl := []int{100, 200}
// 	// sl2 := sl

// 	// sl2[0] = 1000
// 	// // スライスは参照型なので更新された値で表示されてしまう
// 	// fmt.Println(sl)

// 	sl := []int{1, 2, 3, 4, 5}
// 	sl2 := make([]int, 5, 10)
// 	fmt.Println(sl2)

// 	n := copy(sl2, sl)

// 	fmt.Println(n, sl2)
// }

// スライス for

// func main() {
// 	sl := []string{"A", "B", "C"}
// 	fmt.Println(sl)

// 	// for i := range sl {
// 	// 	fmt.Println(i, sl[i])
// 	// }

// 	for i := 0; i < len(sl); i++ {
// 		fmt.Println(i, sl[i])
// 	}
// }

// スライス 可変長引数

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3))
}