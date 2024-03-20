package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

type operationFunc func(a, b int) int

func CalculatePostfix(input string) (string, error) {

	if input == "" {
		return "", fmt.Errorf("invalid expression: empty input")
	}

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
		"^": func(a, b int) int {
			result := 1
			for i := 0; i < b; i++ {
				result *= a
			}
			return result
		},
	}

	tokens := strings.Fields(input)
	for _, token := range tokens {
		value, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, value)
		} else {
			if !isValidOperator(token) {
				return "", fmt.Errorf("invalid expression: unsupported symbol")
			}

			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression: not enough operands")
			}

			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			result := operations[token](a, b)
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid expression: incorrect number of operands")
	}

	resultStr := fmt.Sprintf("%d", stack[0])
	return resultStr, nil
}

func isValidOperator(op string) bool {
	supportedOps := map[string]bool{"+": true, "-": true, "*": true, "/": true, "^": true}
	_, exists := supportedOps[op]
	return exists
}
