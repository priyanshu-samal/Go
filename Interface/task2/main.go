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

func send(n Notifier) {
	fmt.Println(n.Notify())
}

func main() {
	email := Email{address: "user@example.com"}
	send(email)
}
