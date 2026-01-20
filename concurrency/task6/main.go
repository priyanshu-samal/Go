package main

import (
	"fmt"
	"time"
)

func worker(ch chan bool){
	time.Sleep(100*time.Microsecond)
	fmt.Println("Worker done")
	ch<-true
}

func main() {
	ch:=make(chan bool)
	go worker(ch)
	fmt.Println("Waiting")
	<-ch
	
	

	

}