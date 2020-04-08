package queue

type queuer interface {
	Enqueue(n int) error
	Dequeue() (int, error)
}

func NewArrayQueue(size int) queuer {
	arr := make([]int, size)
	return &arrayqueue{1,1, size,&arr}
}

func NewLinkedQueue(size int) queuer {
	tail := new(linkedNode)
	head := &linkedNode{0, nil, tail}
	tail.Prev = head
	return &linkedqueue{head, tail, 0, size}
}