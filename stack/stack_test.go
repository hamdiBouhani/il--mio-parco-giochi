package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(10)
	stack.Pop()
	stack.Push(7)
	stack.Push(3)
	if stack.Peek() != 3 {
		t.Error("Peek() failed")
	}
}
