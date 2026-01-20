package main

import "fmt"

func praducer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go praducer(ch)
	for i := range ch {
		fmt.Println(i)
	}
	<-ch

}