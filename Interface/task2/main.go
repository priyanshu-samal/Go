package main

import "fmt"

type Notifier interface {
	Notify() string
}

type Email struct {
	address string
}

func (e Email) Notify() string {
	return "Email sent to " + e.address
}


//n ye keh rha hai ki mujhe nhi pta notifier kya jhai mujhe ye pta hai ki vo notify frr sakta hai kese vo bhi nhi pta
func send(n Notifier) {
	fmt.Println(n.Notify())
}

func main() {
	email := Email{address: "user@example.com"}
	send(email)
}
