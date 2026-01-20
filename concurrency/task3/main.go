package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		 fmt.Println(i)
	}
}

func printLetters(){
	for ch:='a';ch<='e';ch++{
		time.Sleep(100*time.Millisecond)
		 fmt.Println(string(ch))
	}
}

func main() {
   go printNumbers()
   go printLetters()
   time.Sleep(1* time.Second)
}