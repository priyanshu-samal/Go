package main

import "fmt"

type Stack struct {
	data []int
	top  int
	size int
}

// constructor
func NewStack(size int) *Stack {
	return &Stack{
		data: make([]int, size),
		top:  -1,   // empty stack
		size: size, // capacity
	}
}

// push element onto stack
func (s *Stack) Push(x int) bool {
	if s.top == s.size-1 {
		return false // stack overflow
	}

	s.top++
	s.data[s.top] = x
	return true
}

// pop element from stack
func (s *Stack) Pop() (int, bool) {
	if s.top == -1 {
		return 0, false // stack underflow
	}

	val := s.data[s.top]
	s.top--
	return val, true
}

// peek top element
func (s *Stack) Peek() (int, bool) {
	if s.top == -1 {
		return 0, false
	}
	return s.data[s.top], true
}

// check if empty
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func main() {
	s := NewStack(3)

	fmt.Println(s.Push(10)) // true
	fmt.Println(s.Push(20)) // true
	fmt.Println(s.Push(30)) // true
	fmt.Println(s.Push(40)) // false (overflow)

	fmt.Println(s.Pop()) // 30 true
	fmt.Println(s.Pop()) // 20 true
	fmt.Println(s.Pop()) // 10 true
	fmt.Println(s.Pop()) // 0 false (underflow)
}
