package main

import "fmt"

type RequestCounter struct {
	count int
}

func (c *RequestCounter) hit() {
	c.count += 1
}

func (c *RequestCounter) reset() {
	c.count = 0
}

func (c RequestCounter) value() int {
	return c.count
}

func main() {
	v := RequestCounter{0}
	v.hit()
	v.hit()
	v.hit()
	fmt.Println(v.value())
	v.reset()
	fmt.Println(v.value())
}