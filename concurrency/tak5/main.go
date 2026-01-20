package main

import "fmt"

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go producer(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
