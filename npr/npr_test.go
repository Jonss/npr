package npr

import "testing"

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
