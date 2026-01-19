package main

import "fmt"

type Box[T any] struct {
	value T
}

func main() {
	intBox := Box[int]{value: 10}
	strBox := Box[string]{value: "golang"}

	fmt.Println(intBox.value)
	fmt.Println(strBox.value)
}
