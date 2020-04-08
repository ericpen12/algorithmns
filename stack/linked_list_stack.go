package stack

import "errors"

type linkedNode struct {
	Val  int
	Next *linkedNode
}

type stack struct {
	top  int
	len  int
	cap  int
	head *linkedNode
}

type stacker interface {
	Push(n int) error
	Pop() (int, error)
	IsEmpty() bool
}

func NewStack(size int) stacker {
	return &stack{
		top:  0,
		len:  0,
		cap: size,
		head: new(linkedNode),
	}
}

func (s *stack) Push(n int) error {
	if s.len > s.cap {
		return errors.New("stack overflow")
	}
	s.len++
	next := s.head.Next
	s.head.Next = &linkedNode{n, next}
	return nil
}

func (s *stack) Pop() (int, error) {
	if s.len == 0 {
		return 0, errors.New("stack underflow")
	}
	s.len--
	node := s.head.Next
	s.head.Next = node.Next

	return node.Val, nil
}

func (s *stack) IsEmpty() bool {
	return s.len == 0
}
