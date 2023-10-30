package main

import (
	"errors"
	"fmt"
	"sync"
)

type Stack[T any] struct {
	list []T
	mu   sync.Mutex
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Empty() bool {
	return len(s.list) == 0
}

func (s *Stack[T]) Peek() T {
	if s.Empty() {
		err := errors.New("Stack is empty")
		panic(err)

	}
	return s.list[len(s.list)-1]
}

func (s *Stack[T]) Push(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.list = append(s.list, item)
}

func (s *Stack[T]) Pop() T {
	var elm T
	s.mu.Lock()
	defer s.mu.Unlock()
	elm, s.list = s.Peek(), s.list[0:len(s.list)-1]
	return elm
}

func main() {
	stack := NewStack[int]()
	stack.Push(10)
	stack.Pop()
	stack.Push(7)
	stack.Push(3)
	fmt.Println(stack.Peek())
}
