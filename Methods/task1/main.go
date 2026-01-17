package main

import "fmt"

type side struct {
	l, b int
}

func (s side) area() int {
	return s.l * s.b
}

func (s side) para() int {
	return 2 * (s.l + s.b)
}

func main() {
	s := side{10, 20}
	fmt.Println(s.area(),s.para())
}