package array_stack

import "errors"

type stack struct {
	top   int
	array *[]int
}

type stacker interface {
	Push(n int)
	Pop() (int, error)
	IsEmpty() bool
}

func (s *stack) Push(n int) {
	s.top++
	(*s.array)[s.top] = n
}

func (s *stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("underflow")
	}
	s.top--
	return (*s.array)[s.top+1], nil
}

func (s *stack) IsEmpty() bool {
	return s.top == 0
}

func NewStack(size int) stacker {
	arr := make([]int, size)
	return &stack{
		top:   0,
		array: &arr,
	}
}
