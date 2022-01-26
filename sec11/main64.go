package main

import "fmt"

type T struct {
	// User User
	User
}

type User struct {
	Name string
	Age  int
	// X, Y int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
  t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)

	fmt.Println(t.User)

	fmt.Println(t.User.Name)

	// Tの宣言でUser User（名前と型が一致）ではなくUserで宣言すると、t.Nameでも可能
	fmt.Println(t.Name)

	t.User.SetName()
	fmt.Println(t)
}