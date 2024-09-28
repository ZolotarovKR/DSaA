package operators

import "testing"

func TestMaxValue(t *testing.T) {
	testCases := []struct {
		ns       []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 36},
		{[]int{1, 2, 0, 4}, 12},
		{[]int{1, 0, 3, 4}, 16},
		{[]int{1, 2, 3, 1}, 12},
		{[]int{1, 1, 0, 4}, 8},
	}

	for _, tc := range testCases {
		got := maxValue(tc.ns)
		if got != tc.expected {
			t.Fatalf(
				"Expected MaxValue(%v) to equal %v, but got %v",
				tc.ns,
				tc.expected,
				got,
			)
		}
	}
}
