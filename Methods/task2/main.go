package main

import "fmt"

type Value struct {
	x, y int
}

func (v *Value) change() {
	v.x = v.x * 10
	v.y = v.y * 10
}

func (v Value) area() int {
	return v.x * v.y
}

func main() {
	a := Value{10, 20}

	a.change()                 
	fmt.Println(a.area())      
}
