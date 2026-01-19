package main

import "fmt"

type Counter struct {
	x int
}


func (c *Counter) Inc() {
	c.x++
}

func main() {
	v := Counter{10}

	v.Inc()
	v.Inc()
	v.Inc()

	fmt.Println(v.x) 
}
