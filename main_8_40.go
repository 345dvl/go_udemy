package main

import "fmt"

func main() {
	Loop:
		for {
			for {
				for {
					fmt.Println("START")
					break Loop
				}
				fmt.Println("処理をしない")
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("END")
}