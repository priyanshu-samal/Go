package main

import (
	"fmt"
	"time"
)

func worker1(ch chan bool){
	time.Sleep(100*time.Millisecond)
	fmt.Println("Worker1 completed")
	ch<-true
}

func worker2(ch chan bool){
	time.Sleep(100*time.Millisecond)
	fmt.Println("Worker2 completed")
	ch<-true
}

func main() {
	ch := make(chan bool)
	fmt.Println("Wainting")
	go worker1(ch)
	go worker2(ch)

	<-ch
	<-ch
}