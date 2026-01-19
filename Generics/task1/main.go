package main

import "fmt"

func Print[T any](v T) {
	fmt.Println(v)
}

func main() {
	Print(10)
	Print("golang")
	Print(true)

	type User struct {
		name string
	}

	u := User{name: "Ravi"}
	Print(u)
}
