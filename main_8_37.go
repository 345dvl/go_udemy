package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	for k := 1; k <= 10; k++ {
		if k == 3 { continue }
		fmt.Println(k)
	}

	arr := [3]int{5, 6, 7}
	for i2:=0; i2 < len(arr); i2++ {
		fmt.Println(arr[i2])
	}

	arr2 := [3]int{4, 5, 6}
	for i3, v := range arr2 {
		fmt.Println(i3, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k4, v4 := range m {
		fmt.Println(k4, v4)
	}
}