package dynamicprogramming

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	n := 10 // Change this to the desired Fibonacci number
	result := Fibonacci(n)
	t.Logf("The %dth Fibonacci number is: %d\n", n, result)
}
