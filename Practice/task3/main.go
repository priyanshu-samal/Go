package main

import "fmt"

type Speaker interface {
	Speak()
}

type Human struct {
	x string
}

type Robot struct {
	y string
}

func (h Human) Speak() {
	fmt.Println(h.x)
}

func (r Robot) Speak(){
	fmt.Println(r.y)
}

func Talk(s Speaker){
	s.Speak()
}

func main() {
    v:=Human{"Helooo"}
	y:=Robot{"hiiiii"}

	Talk(v)
	Talk(y)
}