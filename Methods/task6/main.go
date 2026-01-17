package main

import "fmt"

type Counter struct {
	value int
}

func (c *Counter) inc() {
	c.value = c.value + 1
}

func (c *Counter) add(a int) {
	c.value = c.value + a
}

func (c *Counter) get() int {
	return c.value
}

func main() {

	x := Counter{10}
	x.inc()
	x.add(5)
	fmt.Println(x.get())

}