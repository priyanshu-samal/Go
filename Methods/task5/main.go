package main

import "fmt"

type Vertex struct {
	a, b int
}

func (v *Vertex) scale(f int) {
	v.a = v.a * f
	v.b = v.b * f
}

func (v *Vertex) abs() int {
	return v.a * v.b
}

func main() {
	v := Vertex{10, 20}

	v.scale(2)            // command: mutate state
	fmt.Println(v.abs())  // query: read state
}
