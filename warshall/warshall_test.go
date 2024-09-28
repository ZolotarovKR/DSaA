package warshall

import "testing"

func TestWarshall(t *testing.T) {
	testCases := []struct {
		j, i     int
		expected bool
	}{
		{0, 1, true},
		{0, 3, true},
		{1, 1, true},
		{1, 0, true},
		{2, 2, false},
		{2, 1, false},
		{3, 2, true},
		{3, 3, true},
	}
	var g = [][]bool{
		{false, true, false, false},
		{false, false, false, true},
		{false, false, false, false},
		{true, false, true, false},
	}
	w := Warshall(g)
	for _, r := range w {
		t.Log(r)
	}

	for _, tc := range testCases {
		got := w[tc.j][tc.i]
		if got != tc.expected {
			t.Fatalf(
				"Expected Warshall at (%v, %v) to equal %v, but got %v",
				tc.j,
				tc.i,
				tc.expected,
				got,
			)
		}
	}

}
