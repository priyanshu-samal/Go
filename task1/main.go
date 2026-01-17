package main

import "fmt"

type Para struct {
	l, b int
}

func (p Para) parameter() int {
	return 2 * (p.l + p.b)
}

func main() {
	v := Para{20, 30}
	fmt.Println(v.parameter())

}