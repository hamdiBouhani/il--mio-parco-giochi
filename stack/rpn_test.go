package stack

import (
	"fmt"
	"testing"
)

func TestEvaluatePostfixExpression(T *testing.T) {
	expression := "3 4 + 5 *"
	result, err := EvaluatePostfixExpression(expression)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}
