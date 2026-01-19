package main

import "fmt"

func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func main() {
	x := 10
	y := 20
	Swap(&x, &y)
	fmt.Println(x, y)

	s1 := "hello"
	s2 := "world"
	Swap(&s1, &s2)
	fmt.Println(s1, s2)
}
