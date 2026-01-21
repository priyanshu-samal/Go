package main

import "fmt"

func main() {
	var v interface{} = "hello"

	s, ok := v.(string)

	fmt.Println(s, ok)
}
