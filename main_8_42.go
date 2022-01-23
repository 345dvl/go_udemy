package main

import "fmt"

func main() {
	// panicはPGMを強制的にエラーとさせ終了させるため、使用は推奨されない
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("start")
}