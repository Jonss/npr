package npr

import (
	"errors"
	"fmt"
	"math"
)

type operation func(a, b int) (int, error)

var symbolMap = map[string]operation{
	"+": func(a, b int) (int, error) {
		return a + b, nil
	},
	"-": func(a, b int) (int, error) {
		return a - b, nil
	},
	"*": func(a, b int) (int, error) {
		return a * b, nil
	},
	"/": func(a, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("error division by zero")
		}
		return a / b, nil
	},
	"sqrt": func(a, b int) (int, error) {
		fmt.Println("sqrt", a, b)
		s := int(math.Sqrt(float64(a)))
		fmt.Println(s)
		return s, nil
	},
}

func function(s string) (operation, error) {
	operation, ok := symbolMap[s]
	if !ok {
		return func(a, b int) (int, error) { return 0, nil }, errors.New("unknown token")
	}

	return operation, nil
}
