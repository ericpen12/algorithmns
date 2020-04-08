package array_queue

import "errors"

type queue struct {
	head  int
	tail  int
	len   int
	array *[]int
}

type queuer interface {
	Enqueue(n int) error
	Dequeue() (int, error)
}

func (q *queue) Enqueue(n int) error {
	if q.tail + 1 == q.head {
		return errors.New("queue overflow")
	}
	(*q.array)[q.tail] = n
	q.tail++
	if q.tail == q.len {
		q.tail = 0
	}
	return nil
}

func (q *queue) Dequeue() (int, error) {
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

func NewQueue(size int) queuer {
	arr := make([]int, size)
	return &queue{
		head:  1,
		tail:  1,
		len:   size,
		array: &arr,
	}
}
