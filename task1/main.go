package main

import "fmt"

type mul struct {
	x, y int
}

func (m mul) krunga() int {
	return m.x * m.y
}

func main() {
	val := mul{20, 30}
	fmt.Println(val.krunga())
}