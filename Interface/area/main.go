package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct {
	l, b float64
}

type Circle struct {
	r float64
}

func (s Square) Area() float64 {
	return s.l * s.b
}

func (c Circle) Area() float64 {
	return 3.14 * c.r * c.r
}

func printArea(z Shape) float64{
	return z.Area()
}


func main(){
	x:=Square{10,20}
	y:=Circle{10}
	fmt.Println(printArea(x))
	fmt.Println(printArea(y))

}