package main

import (
	"errors"
	"fmt"
)

func getAge(m map[string]int, name string) (int, error) {
	if m == nil {
		return 0, errors.New("map is nil")
	}

	age, ok := m[name]
	if !ok {
		return 0, errors.New("name not found")
	}

	return age, nil
}

func main() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	age, err := getAge(ages, "Alice")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("age:", age)
}
