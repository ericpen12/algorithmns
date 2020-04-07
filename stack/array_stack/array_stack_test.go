package array_stack_test

import (
	"datastures/stack/array_stack"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := array_stack.NewStack(20)
	stack.Push(10)
	stack.Push(120)
	stack.Push(101)

	for !stack.IsEmpty() {
		v, _:= stack.Pop()
		t.Log(v)
	}
}

func TestStackOverflow(t *testing.T) {
	stack := array_stack.NewStack(2)
	stack.Push(10)
	stack.Push(120)
	err := stack.Push(101)
	if err == nil || err.Error() != "stack overflow"{
		t.Error("without overflow")
	}
	t.Log(err)
}

func TestStackUnderflow(t *testing.T) {
	stack := array_stack.NewStack(2)
	stack.Push(10)

	stack.Pop()
	_, err := stack.Pop()
	if err == nil || err.Error() != "stack underflow" {
		t.Error("without underflow")
	}
	t.Log(err)
}


