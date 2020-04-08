package stack

import "errors"

type linkedNode struct {
	Val  int
	Next *linkedNode
}

type linkedstack struct {
	top  int
	len  int
	cap  int
	head *linkedNode
}

func (s *linkedstack) Push(n int) error {
	if s.len > s.cap {
		return errors.New("stack overflow")
	}
	s.len++
	next := s.head.Next
	s.head.Next = &linkedNode{n, next}
	return nil
}

func (s *linkedstack) Pop() (int, error) {
	if s.len == 0 {
		return 0, errors.New("stack underflow")
	}
	s.len--
	node := s.head.Next
	s.head.Next = node.Next

	return node.Val, nil
}

func (s *linkedstack) IsEmpty() bool {
	return s.len == 0
}
