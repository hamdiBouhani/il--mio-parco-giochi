package main

import (
	"fmt"
)

func main() {
	stack := NewStack[int]()
	stack.Push(10)
	stack.Pop()
	stack.Push(7)
	stack.Push(3)
	fmt.Println(stack.Peek())
}
