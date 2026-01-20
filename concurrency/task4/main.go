package main

import (
	"fmt"
	"time"
)

func numbers(ch chan bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(i)
	}
	ch<- true
}

func character(ch chan bool){
	for i:='a';i<='e';i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(string(i))
	}
	ch<-true
}

func main() {
     ch:=make(chan bool, 1)
	 go numbers(ch)
	 go character(ch)
	 <-ch
	 <-ch

}