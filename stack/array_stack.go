package stack

import "errors"

type arraystack struct {
	top   int
	len   int
	array *[]int
}

func (s *arraystack) Push(n int) error {
	s.top++

	if s.top >= s.len {
		return errors.New("stack overflow")
	}

	(*s.array)[s.top] = n
	return nil
}

func (s *arraystack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack underflow")
	}
	s.top--
	return (*s.array)[s.top+1], nil
}

func (s *arraystack) IsEmpty() bool {
	return s.top == 0
}
