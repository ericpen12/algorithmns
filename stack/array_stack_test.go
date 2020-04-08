package stack_test

import (
	stack2 "datastures/stack"
	"testing"
)

func TestNewArrayStack(t *testing.T) {
	s := stack2.NewArrayStack(20)
	s.Push(10)
	s.Push(120)
	s.Push(101)

	for !s.IsEmpty() {
		v, _:= s.Pop()
		t.Log(v)
	}
}

func TestStackOverflow(t *testing.T) {
	stack := stack2.NewArrayStack(2)
	stack.Push(10)
	stack.Push(120)
	err := stack.Push(101)
	if err == nil || err.Error() != "stack overflow"{
		t.Error("without overflow")
	}
	t.Log(err)
}

func TestStackUnderflow(t *testing.T) {
	stack := stack2.NewArrayStack(2)
	stack.Push(10)

	stack.Pop()
	_, err := stack.Pop()
	if err == nil || err.Error() != "stack underflow" {
		t.Error("without underflow")
	}
	t.Log(err)
}


