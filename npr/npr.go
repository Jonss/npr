package npr

import (
	"strconv"
	"strings"
)

func NPRCalc(expression string) (int64, error) {
	nums := strings.Split(expression, " ")
	stack := NewStack()
	result := 0

	for _, value := range nums {
		n, err := strconv.Atoi(value)
		if err != nil {
			operation, err := function(value)
			if err != nil {
				return 0, err
			}
			v2, v1 := 0, 0
			if value == "sqrt" {
				v2, v1 = 0, stack.Pop()
			} else {
				v2, v1 = stack.Pop(), stack.Pop()
			}

			result, err = operation(v1, v2)
			if err != nil {
				return 0, err
			}
			stack.Push(result)
		}
		if n != 0 {
			stack.Push(n)
		}
	}

	return int64(result), nil
}
