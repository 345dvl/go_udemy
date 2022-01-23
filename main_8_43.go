package main

import (
	"fmt"
)

// go goroutin スレッドよりも小さい単位で行う並行処理

// func sub() {
// 	for {
// 		fmt.Println(("Sub Loop"))
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func main() {
// 	go sub()
// 	go sub()

// 	for {
// 		fmt.Println("Main Loop")
// 		time.Sleep(200 * time.Millisecond)
// 	}
// }


// init

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}