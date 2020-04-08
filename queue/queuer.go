package queue

type queuer interface {
	Enqueue(n int) error
	Dequeue() (int, error)
}

func NewArrayQueue(size int) queuer {
	arr := make([]int, size)
	return &arrayqueue{
		head:  1,
		tail:  1,
		len:   size,
		array: &arr,
	}
}