package lis

import (
	"slices"
	"testing"
)

func TestLis(t *testing.T) {
	testCases := []struct{ xs, expected []int }{
		{[]int{50, 3, 10, 7, 40, 80, 14, 19, 99}, []int{3, 7, 14, 19, 99}},
		{[]int{10, 22, 11, 0, 1, 18, 0, 3}, []int{10, 11, 18}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		got := Lis(tc.xs)
		if len(got) != len(tc.expected) || !slices.IsSorted(got) {
			t.Fatalf(
				"Expected LIS of %v to be %v, but got %v",
				tc.xs,
				tc.expected,
				got,
			)
		}
	}
}
