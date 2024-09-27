package shortest_path

import "testing"

func TestShortestPath(t *testing.T) {
	testCases := []struct {
		from, to int
		expected float64
	}{
		{1, 3, 5},
		{5, 3, 4},
		{2, 3, 5},
		{0, 7, 3},
		{4, 3, 2},
	}

	for _, tc := range testCases {
		got := ShortestPath(tc.from, tc.to)
		if got != tc.expected {
			t.Log(topologicalSort())
			t.Fatalf(
				"Expected ShortestPath(%v, %v) to equal %v, but got %v",
				tc.from,
				tc.to,
				tc.expected,
				got,
			)
		}
	}
}
