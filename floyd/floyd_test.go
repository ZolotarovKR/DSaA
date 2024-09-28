package floyd

import (
	"math"
	"testing"
)

func TestFloyd(t *testing.T) {
	testCases := []struct {
		j, i     int
		expected float64
	}{
		{0, 1, 10},
		{0, 3, 4},
		{1, 1, 0},
		{1, 0, 2},
		{2, 2, 0},
		{2, 1, 7},
		{3, 2, 9},
		{3, 3, 0},
	}
	i := math.Inf(1)
	g := [][]float64{
		{0, i, 3, i},
		{2, 0, i, 1},
		{i, 7, 0, 1},
		{6, i, i, 0},
	}
	f := Floyd(g)

	for _, tc := range testCases {
		got := f[tc.j][tc.i]
		if got != tc.expected {
			t.Fatalf(
				"Expected Floyd at (%v, %v) to equal %v, but got %v",
				tc.j,
				tc.i,
				tc.expected,
				got,
			)
		}
	}
}
