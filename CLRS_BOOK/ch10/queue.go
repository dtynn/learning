package ch10

import "fmt"

func NewQueue(capacity int) *Queue {
	return &Queue{
		queue: make([]int, capacity),
	}
}

type Queue struct {
	queue      []int
	head, tail int
}

func (q *Queue) Enqueue(x int) error {
	if q.tail-q.head >= len(q.queue) {
		return fmt.Errorf("queue overflow")
	}

	q.queue[(q.tail)%len(q.queue)] = x
	q.tail++
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.tail == q.head {
		return NIL, fmt.Errorf("queue underflow")
	}

	x := q.queue[q.head%len(q.queue)]
	q.head++
	return x, nil
}
