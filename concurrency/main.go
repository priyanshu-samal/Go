package main

import (
	"fmt"
	"sync"
	"time"
)

// worker function to simulate fetching a URL
func worker(wg *sync.WaitGroup, resultChan chan string, jobsChan chan string) {
	defer wg.Done() // Ensure that Done is called when the function exits

	for url := range jobsChan {
		// Simulate network delay
		time.Sleep(time.Millisecond * 50)
		fmt.Println("Fetching URL:", url)
		resultChan <- "Fetched " + url
	}
}

func main() {
	jobs := []string{
		"http://example.com/image1.jpg",
		"http://example.com/image2.jpg",
		"http://example.com/image3.jpg",
		"http://example.com/image4.jpg",
		"http://example.com/image5.jpg",
	}

	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish
	totalWorkers := 5      // Total number of workers
	resultChan := make(chan string, len(jobs))
	jobsChan := make(chan string, len(jobs))

	start := time.Now()

	// Start workers
	for i := 0; i < totalWorkers; i++ {
		wg.Add(1)
		go worker(&wg, resultChan, jobsChan)
	}

	// Send jobs to workers
	for _, job := range jobs {
		jobsChan <- job
	}
	close(jobsChan) // Close the jobs channel to signal that no more jobs will be sent

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(resultChan) // Close the result channel after all workers are done
	}()

	// Collect results
	for result := range resultChan {
		fmt.Println("Result received:", result)
	}

	fmt.Println("Total time taken:", time.Since(start))
}
