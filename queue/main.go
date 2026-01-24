package main

import "fmt"

type Queue struct {
	rear  int
	front int
	size  int
	data  []int
}

func NewQueue(size int) *Queue {
	return &Queue{
		data:  make([]int, size),
		size:  size,
		front: -1,
		rear:  -1,
	}
}

func (q *Queue) enque(x int) bool {
	if q.rear == q.size-1 {
		return false
	}
	if q.front == -1 {
		q.front = 0
		q.rear = 0
		q.data[q.rear] = x
		q.rear++
		return true
	}
	q.rear = x
	q.rear++
	return true
}

func (q *Queue) deque() (int, bool) {
	if q.rear == -1 {
		return 0, false
	}
	if q.front == q.size-1 {
		x := q.data[q.front]
		q.front = -1
		q.rear = -1
		return x, true
	}
	if q.front == q.rear {
		x := q.data[q.front]
		q.front = -1
		q.rear = -1
		return x, true
	}
	x := q.data[q.front]
	q.front++
	return x, true
}

func main() {
	q := NewQueue(5)
	q.enque(10)
	q.enque(13)
	q.enque(11)
	q.enque(18)
	fmt.Println(q.deque())
	q.enque(11)
	fmt.Println(q.data)
}