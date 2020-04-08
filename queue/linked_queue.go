package queue

import "errors"

type linkedNode struct {
	Val  int
	Prev *linkedNode
	Next *linkedNode
}

type linkedqueue struct {
	head *linkedNode
	tail *linkedNode
	len  int
	cap  int
}

func (q *linkedqueue) Enqueue(n int) error {
	if q.len == q.cap {
		return errors.New("queue overflow")
	}
	q.len++
	prev := q.tail.Prev
	node := &linkedNode{n, prev, q.tail}
	prev.Next = node
	q.tail.Prev = node

	return nil
}

func (q *linkedqueue) Dequeue() (int, error) {
	if q.len == 0 {
		return 0, errors.New("queue underflow")
	}
	q.len--
	next := q.head.Next
	q.head.Next = next.Next
	next.Next.Prev = q.head

	return next.Val, nil
}
