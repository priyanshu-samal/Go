package main

import (
	"fmt"
	"time"
)

func signal(ch chan bool) {
	time.Sleep(300*time.Millisecond)
	fmt.Println("done")
	ch<-true
}

func main() {
	ch := make(chan bool)
	go signal(ch)
	<-ch
}