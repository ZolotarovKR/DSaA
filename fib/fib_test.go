package fib

import "testing"

func TestFib(t *testing.T) {
	testCases := []struct{ n, expected int }{
		{10, 55},
		{40, 102334155},
		{1, 1},
		{20, 6765},
		{12, 144},
		{46, 1836311903},
		{0, 0},
	}

	for _, tc := range testCases {
		got := Fib(tc.n)
		if got != tc.expected {
			t.Fatalf(
				"Expected Fib(%v) to equal %v, but got %v",
				tc.n,
				tc.expected,
				got,
			)
		}
	}
}
