package main

import "fmt"

type Account struct {
	balance int
}

func (a *Account) deposit(money int) {
	a.balance += money
}

func (a *Account) withdrawl(money int) {
	a.balance -= money
	if a.balance < 0 {
		a.balance = 0
	}
}

func (a Account) getBalance() int {
	return a.balance
}

func main() {
	v := Account{100}
	fmt.Println(v.getBalance())
	v.deposit(100)
	fmt.Println(v.getBalance())
	v.withdrawl(100)
	fmt.Println(v.getBalance())


}