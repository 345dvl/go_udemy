package main

import "fmt"

// ポインタ

// int a → 数値
// &a → aのアドレス
// int *p → ポインタ変数、pにはp = &aのようにアドレスをセット可能
// fmt.Println(p)  → aのアドレス
// fmt.Println(*p) → aの中身

func Double(i int) {
	i = i * 2
}

func Doublev2(i *int) {
	*i = *i * 2
}

func Doublev3(s []int) {
	for i, v := range s  {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n) //100

	fmt.Println(&n) //0x...

	Double(n)

	fmt.Println(n) //100

	var p *int = &n
	fmt.Println(p) //0x...
	fmt.Println(*p) //100

	*p = 300
	fmt.Println(n)
	n = 200
	fmt.Println(*p)

	Doublev2(&n)
	fmt.Println(n)

	Doublev2(p)
	fmt.Println(*p)

	var sl []int = []int{1,2,3}

	Doublev3(sl)
	fmt.Println(sl)
}