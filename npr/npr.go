package npr

import (
	"strconv"
	"strings"
)

func NPRCalc(expression string) (int64, error) {
	nums := strings.Split(expression, " ")
	stack := NewStack()

	result := 0
	for _, i := range nums {
		n, err := strconv.Atoi(i)
		if err != nil {
			operation, err := validSymbol(i)
			if err != nil {
				return 0, err
			}
			v2, v1 := stack.Pop(), stack.Pop()

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
