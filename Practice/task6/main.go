package main

import "fmt"

type count struct {
	x int
}

func (c count) value() int {
	c.x++
	return c.x
}

func (c *count) point() int {
	c.x++
	return c.x
}

func main() {
	v := count{10}
	fmt.Println(v.value())
	fmt.Println(v.point())
}