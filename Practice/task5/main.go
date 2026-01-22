package main

import "fmt"

func main() {
	s := make([]int, 0)
	for i := 0; i <= 10; i++ {
		s = append(s, i)
	}
	fmt.Println(s)

}