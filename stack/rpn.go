package stack

import (
	"fmt"
	"strconv"
	"strings"
)

// Reverse Polish Notation (RPN)
// https://en.wikipedia.org/wiki/Reverse_Polish_notation

func EvaluatePostfixExpression(expression string) (int, error) {
	stack := NewStack[int]()

	tokens := strings.Fields(expression)

	for _, token := range tokens {

		if isOperator(token) {

			if stack.Count() < 2 {
				return 0, fmt.Errorf("Not enough operands for operator: %s", token)
			}

			operand2 := stack.Pop()
			operand1 := stack.Pop()

			result, err := applyOperator(token, operand1, operand2)
			if err != nil {
				return 0, err
			}
			stack.Push(result)

		} else {
			operand, err := strconv.Atoi(token)
			if err != nil {
				return 0, fmt.Errorf("Invalid operand: %s", token)
			}
			stack.Push(operand)
		}
	}
	return stack.Pop(), nil
}

func isOperator(token string) bool {
	operators := "+-*/"
	return strings.Contains(operators, token)
}

func applyOperator(operator string, operand1, operand2 int) (int, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return operand1 / operand2, nil
	default:
		return 0, fmt.Errorf("Invalid operator: %s", operator)
	}
}
