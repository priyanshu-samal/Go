package main

import "fmt"

type Read struct {
	x int
}

func (r *Read) inc() {
	r.x++
}

func main() {
	v := Read{10}

	v.inc()
	fmt.Println(v);
	
	(&v).inc()
	fmt.Println(v)

}