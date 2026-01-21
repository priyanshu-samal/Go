package main

import (
	"fmt"
	"time"
)

func worker(jobs chan int,done chan struct{}){
	for{
		select{
			case job,ok :=<-jobs:
			if !ok{
				fmt.Println("worker: jobs finished")
				return
			}
			fmt.Println("worker: processing job", job)

			case <-done:
			// shutdown signal
			fmt.Println("worker: shutting down")
			return
		}
	}
}


func main(){

    jobs:=make(chan int)
	done:=make(chan struct{})

	go worker(jobs,done)


	



}