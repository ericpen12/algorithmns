package stack

type stacker interface {
	Push(n int) error
	Pop() (int, error)
	IsEmpty() bool
}

func NewArrayStack(size int) stacker {
	arr := make([]int, size)
	return &arraystack{
		top:   0,
		len:   size,
		array: &arr,
	}
}

func NewLinkStack(size int) stacker {
	return &linkedstack{
		top:  0,
		len:  0,
		cap: size,
		head: new(linkedNode),
	}
}
