package main

import (
	"fmt"
	"time"
)

func producer(ch chan int){
     for i:=1;i<=3;i++{
		fmt.Println("Sending x")
		ch<-i
	 }
	 close(ch)
}

func main() {
	ch := make(chan int,2)
	go producer(ch)
	time.Sleep(100*time.Millisecond)
	for i:=range ch{
		fmt.Println(i)
	}

}