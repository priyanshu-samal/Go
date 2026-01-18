package main

type Printer interface {
	Print() string
}

type User struct {
	name string
}

func (u User) Print() string {
	return u.name
}



func main(){

}