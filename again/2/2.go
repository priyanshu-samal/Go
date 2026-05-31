package main

import "fmt"

func Hello()string{
	return "Hellooo"
}

func main() {
	v := 10
	fmt.Println(v);

	x:=Hello()
	fmt.Println(x)
}