package rod_cutting

import "testing"

func TestCutRod(t *testing.T) {
	prices := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	testCases := []struct {
		length   int
		expected int
	}{
		{4, 10},
		{6, 17},
		{10, 30},
		{0, 0},
		{7, 18},
		{1, 1},
	}

	for _, tc := range testCases {
		got := CutRod(tc.length, prices)
		if got != tc.expected {
			t.Fatalf(
				"Expected CutRod(%v) to equal %v, but got %v",
				tc.length,
				tc.expected,
				got,
			)
		}
	}
}
