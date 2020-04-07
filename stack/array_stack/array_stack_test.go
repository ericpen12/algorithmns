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


