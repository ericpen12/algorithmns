package array_queue

type queue struct {
	head  int
	tail  int
	len   int
	array *[]int
}

type queuer interface {
	Enqueue(n int)
	Dequeue() int
}

func (q *queue) Enqueue(n int) {
	(*q.array)[q.tail] = n
	if q.tail == q.len {
		q.tail = 0
	} else {
		q.tail++
	}
}

func (q *queue) Dequeue() int {
	x := (*q.array)[q.head]

	if q.head == q.len {
		q.head = 0
	} else {
		q.head++
	}

	return x
}

func NewQueue(size int) queuer {
	arr := make([]int, size)
	return &queue{
		head:  0,
		tail:  0,
		len:   size,
		array: &arr,
	}
}
