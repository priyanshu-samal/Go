package main

import "fmt"
//Interfaces let you change WHAT you use without changing WHERE you use it.

type Describer interface {
	Describe() string
}

type Book struct {
	title string
}

type Movie struct{
	name string
}

func (b Book) Describe() string {
	return b.title
}

func (m Movie) Describe()string{
	return m.name
}

func tellMe(n Describer) {
	fmt.Println(n.Describe())
}



func main() {

	v := Book{title: "wherever"}
	x:= Movie{name:"till"}
    tellMe(v)
	tellMe(x)
}