package main

import (
	"errors"
	"fmt"
	"os"
)

// One of the most important aspects of Go is that errors are values.
// They are not exceptions. You don't "throw" them; you return them.

// 1. Defining a Custom Error
// We can define our own error types to carry more context.
type ConnectionError struct {
	URL      string
	Attempts int
}

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("failed to connect to %s after %d attempts", e.URL, e.Attempts)
}

// 2. Sentinel Errors
// These are static error values used for comparison (errors.Is).
var ErrNotFound = errors.New("item not found")

func findItem(id int) (string, error) {
	if id == 42 {
		return "Everything", nil
	}
	// Return the sentinel error
	return "", ErrNotFound
}

func connectToService(url string) error {
	// Simulate a failure
	return &ConnectionError{URL: url, Attempts: 3}
}

func main() {
	// Example 1: Standard Error Handling
	// We call a function and immediately check if err != nil.
	item, err := findItem(100)
	if err != nil {
		// usage of errors.Is to check for specific sentinel errors
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Search failed: The item does not exist.")
		} else {
			fmt.Println("Search failed: Unknown error.")
		}
	} else {
		fmt.Println("Found:", item)
	}

	// Example 2: Custom Error Types
	// usage of errors.As to unwrap and inspect the error struct
	err = connectToService("http://example.com")
	if err != nil {
		var connErr *ConnectionError
		if errors.As(err, &connErr) {
			fmt.Printf("Connection Error Details -> URL: %s | API Key: Refred to retry %d times\n", connErr.URL, connErr.Attempts)
		} else {
			fmt.Println("Generic error:", err)
		}
	}
}
