package array_stack

import "errors"

type stack struct {
	top   int
	len   int
	array *[]int
}

type stacker interface {
	Push(n int) error
	Pop() (int, error)
	IsEmpty() bool
}

func (s *stack) Push(n int) error {
	s.top++

	if s.top >= s.len {
		return errors.New("stack overflow")
	}

	(*s.array)[s.top] = n
	return nil
}

func (s *stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack underflow")
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
		len:   size,
		array: &arr,
	}
}
