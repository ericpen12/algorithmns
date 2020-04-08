package stack_test

import (
	"datastures/stack"
	"testing"
)

func TestNewLinkStack(t *testing.T) {
	s := stack.NewLinkStack(20)
	s.Push(10)
	s.Push(120)
	s.Push(101)

	for !s.IsEmpty() {
		v, _:= s.Pop()
		t.Log(v)
	}
}

func TestLinkStackOverflow(t *testing.T) {
	s := stack.NewLinkStack(2)
	s.Push(10)
	s.Push(120)
	err := s.Push(101)
	if err == nil || err.Error() != "stack overflow"{
		t.Error("without overflow")
	}
	t.Log(err)
}

func TestLinkedStackUnderflow(t *testing.T) {
	s := stack.NewLinkStack(2)
	s.Push(10)

	s.Pop()
	_, err := s.Pop()
	if err == nil || err.Error() != "stack underflow" {
		t.Error("without underflow")
	}
	t.Log(err)
}


