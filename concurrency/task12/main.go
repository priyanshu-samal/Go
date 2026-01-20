package main

import (
	"fmt"
	"time"
)

func first(ch1 chan string) {
	time.Sleep(300*time.Millisecond)
	ch1<-"from ch1"
}

func second(ch2 chan string) {
	time.Sleep(600*time.Millisecond)
	ch2<-"from ch2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go first(ch1)
	go second(ch2)

     select{
		case <-ch1:
		fmt.Println("recived from ch1")

		case <-ch2:
		fmt.Println("recieved from ch2")
		
		case <-time.After(400*time.Millisecond):
		fmt.Println("TImeout")

	 }

}