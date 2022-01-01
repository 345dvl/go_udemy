package main

import (
	"fmt"
	"strconv"
)

// 型

func main() {
	// 整数型
	var i int64 = 100

	fmt.Println(i)
	fmt.Printf("%T\n", i)


	// 文字列型
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	test
	    test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))


	// byte(unit8)型
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))


	// 配列型
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	// インターフェース
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = "Go"
	fmt.Println(x)

	// 型変換
	var j int = 1
	fl64 := float64(j)
	fmt.Println(fl64)
	fmt.Printf("j = %T\n", j)
	fmt.Printf("fl64 = %T\n", fl64)

	j2 := int(fl64)
	fmt.Printf("j2 = %T\n", j2)

	fl := 10.5
	j3 := int(fl)
	fmt.Printf("j3 = %T\n", j3)
	fmt.Println(j3)

	var s1 string = "100"
	fmt.Printf("s1 = %T\n", s1)
	k, _ := strconv.Atoi(s1)
	fmt.Println(k)
	fmt.Printf("k = %T\n", k)

	var i3 int = 200
	s3 := strconv.Itoa(i3)
	fmt.Println(s3)
	fmt.Printf("s3 = %T\n", s3)

	var h string = "Hello World"
	b1 := []byte(h)
	fmt.Println(b1)

	h2 := string(b1)
	fmt.Println(h2)
}