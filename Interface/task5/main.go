package main

import "fmt"

/*
STEP 1: Define the behavior you care about.
This is the interface.
*/
type Notifier interface {
	Notify() string
}

/*
STEP 2: Create a concrete implementation (Email).
*/
type Email struct {
	address string
}

func (e Email) Notify() string {
	return "Email sent to " + e.address
}

/*
STEP 3: Create another implementation (SMS).
*/
type SMS struct {
	number string
}

func (s SMS) Notify() string {
	return "SMS sent to " + s.number
}

/*
STEP 4: Business logic that depends on the INTERFACE,
not on Email or SMS.
*/
func sendNotification(n Notifier) {
	// This function has NO IDEA what concrete type n is.
	// It only knows one thing: n can Notify().
	fmt.Println(n.Notify())
}

func main() {
	email := Email{address: "user@example.com"}
	sms := SMS{number: "9999999999"}

	sendNotification(email)
	sendNotification(sms)
}
