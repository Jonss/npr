package npr

import (
	"errors"
	"strconv"
	"strings"
)

func NPRCalc(expression string) (int64, error) {
	nums := strings.Split(expression, " ")
	stack := []int{}

	symbolMap := map[string]operation{
		"+": func(a, b int) int {
			return a + b
		},
		"-": func(a, b int) int {
			return a - b
		},
		"*": func(a, b int) int {
			return a * b
		},
		"/": func(a, b int) int {
			return a / b
		},
	}

	response := 0
	for idx, i := range nums {
		n, err := strconv.Atoi(i)
		if err != nil {
			symbol, err := validSymbol(i)
			if err != nil {
				return 0, err
			}
			numbers := stack[idx-2:]
			res := symbolMap[symbol]

			response = res(numbers[0], numbers[1])
		}
		if n != 0 {
			stack = append(stack, n)
		}
	}

	return int64(response), nil
}

type operation func(a, b int) int

func validSymbol(s string) (string, error) {
	if s == "+" {
		return "+", nil
	}
	if s == "-" {
		return "-", nil
	}
	if s == "*" {
		return "*", nil
	}
	if s == "/" {
		return "/", nil
	}
	return "", errors.New("unknown token")
}
