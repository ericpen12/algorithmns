package queue

import "errors"

type arrayqueue struct {
	head  int
	tail  int
	len   int
	array *[]int
}

func (q *arrayqueue) Enqueue(n int) error {
	if q.tail+1 == q.head {
		return errors.New("queue overflow")
	}
	(*q.array)[q.tail] = n
	q.tail++
	if q.tail == q.len {
		q.tail = 0
	}
	return nil
}

func (q *arrayqueue) Dequeue() (int, error) {
	if q.tail == q.head {
		return 0, errors.New("queue underflow")
	}

	x := (*q.array)[q.head]
	q.head++
	if q.head == q.len {
		q.head = 0
	}

	return x, nil
}
