package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Human struct {
	name string
}

func (h Human) Speak() string {
	return "Hello, my name is " + h.name
}

func talk(s Speaker) {
	fmt.Println(s.Speak())
}


func main() {
	h := Human{name: "Ravi"}
	talk(h)


}
