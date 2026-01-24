package main

import "fmt"

type Queue struct {
	front int
	rear  int
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

func (q *Queue) Enqueue(x int) bool {
	if q.rear == q.size-1 {
		return false // queue full
	}

	if q.front == -1 {
		q.front = 0
	}

	q.rear++
	q.data[q.rear] = x
	return true
}

func (q *Queue) Dequeue() (int, bool) {
	if q.front == -1 {
		return 0, false // queue empty
	}

	x := q.data[q.front]

	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front++
	}

	return x, true
}

func main() {
	q := NewQueue(5)

	q.Enqueue(10)
	q.Enqueue(13)
	q.Enqueue(11)
	q.Enqueue(18)

	fmt.Println(q.Dequeue()) // 10

	q.Enqueue(99) // still fails after rear hits end

	fmt.Println(q.data)
}
