package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

type operationFunc func(a, b int) int

func CalculatePostfix(input string) (string, error) {
	var stack []int
	operations := map[string]operationFunc{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a / b
		},
	}

	tokens := strings.Fields(input)
	for _, token := range tokens {
		value, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, value)
		} else {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression: not enough operands")
			}
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			opFunc, exists := operations[token]
			if !exists {
				return "", fmt.Errorf("unsupported operator: %s", token)
			}
			result := opFunc(a, b)
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid expression: too many operands")
	}

	resultStr := fmt.Sprintf("%d", stack[0])
	return resultStr, nil
}
