package npr

import "errors"

type operation func(a, b int) (int, error)

func validSymbol(s string) (operation, error) {
	symbolMap := map[string]operation{
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
	}
	fun, ok := symbolMap[s]
	if !ok {
		return func(a, b int) (int, error) { return 0, nil }, errors.New("unknown token")
	}

	return fun, nil
}
