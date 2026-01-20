package main

import "fmt"

func praducer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go praducer(ch)

	for v := range ch {
		if v%2 == 0 {
			fmt.Println(v)
		}
	}

}