package main

import (
	"fmt"
	"log"
	"os"
)

// os

func main() {
	// Exit
	// os.Exit(1)
	// fmt.Println(("Start"))

	// defer実行されない
	// defer func() {
	// 	fmt.Println("defer")
	// }()
	// os.Exit(0)

	// log.Fatal
	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err)
	}
}