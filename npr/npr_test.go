package npr

import (
	"errors"
	"testing"
)

func TestNPRCalc(t *testing.T) {
	testCases := []struct {
		expression string
		want       int64
	}{
		{
			expression: "3 4 +",
			want:       7,
		},
		{
			expression: "10 4 -",
			want:       6,
		},
		{
			expression: "6 7 *",
			want:       42,
		},
		{
			expression: "8 2 /",
			want:       4,
		},
		{
			expression: "5 1 2 + 4 * + 3 -",
			want:       14,
		},
		{
			expression: "9 sqrt",
			want:       3,
		},
		{
			expression: "9 sqrt 3 /",
			want:       1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.expression, func(t *testing.T) {
			got, _ := NPRCalc(tc.expression)
			if got != tc.want {
				t.Fatalf("NPRCalc(). want %v, got %v", tc.want, got)
			}
		})
	}
}

func TestNPRCalc_Error(t *testing.T) {
	testCases := []struct {
		expression string
		err        error
	}{
		{
			expression: "+",
			err:        errors.New("Error: not enough elements for operation"),
		},
		{
			expression: "3 -",
			err:        errors.New("Error: not enough elements for operation"),
		},
		{
			expression: "sqrt",
			err:        errors.New("Error: not enough elements for operation"),
		},
		{
			expression: "10 0 /",
			err:        errors.New("Error: division by zero"),
		},
		{
			expression: "-4 sqrt",
			err:        errors.New("Error: square root of negative number"),
		},
		{
			expression: "3 3",
			err:        errors.New("Error: the user input does not form a valid RPN expression"),
		},
		{
			expression: "abc",
			err:        errors.New("Error: unknown token"),
		},
		{
			expression: "4 5 &",
			err:        errors.New("Error: unknown token/operator"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.expression, func(t *testing.T) {
			_, err := NPRCalc(tc.expression)
			if err != tc.err {
				t.Fatalf("NPRCalc(). want %v, got %v", tc.err, err)
			}
		})
	}
}
