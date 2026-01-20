package main

import (
	"fmt"
	"time"
)

func worker(jobs <-chan int, done <-chan struct{}) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				// jobs channel closed → no more work
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

func main() {
	jobs := make(chan int)
	done := make(chan struct{})

	// start worker
	go worker(jobs, done)

	// producer: send jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
		time.Sleep(300 * time.Millisecond)
	}

	// signal no more jobs
	close(jobs)

	// give worker time to finish (simulating app lifetime)
	time.Sleep(1 * time.Second)

	// shutdown signal (safe even if worker already exited)
	close(done)

	fmt.Println("main: exit")
}
