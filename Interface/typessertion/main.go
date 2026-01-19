package main

import "fmt"

// This function accepts ANY type
func handleValue(v interface{}) {

	fmt.Println("Received value:", v)

	// SAFE type assertion
	if n, ok := v.(int); ok {
		fmt.Println("It is an int, int + 10 =", n+10)
		return
	}

	if s, ok := v.(string); ok {
		fmt.Println("It is a string, length =", len(s))
		return
	}

	fmt.Println("Unknown type")
}

func main() {

	// interface{} holding different types
	var a interface{} = 10
	var b interface{} = "golang"
	var c interface{} = true

	handleValue(a)
	handleValue(b)
	handleValue(c)

	// -------- TYPE CONVERSION (different thing) --------
	var x int = 5
	var y float64 = float64(x)
	fmt.Println("Converted int to float:", y)
}
