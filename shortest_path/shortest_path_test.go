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
	g := Graph{
		{{to: 4, weight: 2}},
		{{to: 0, weight: 1}, {to: 2, weight: 1}, {to: 4, weight: 4}, {to: 5, weight: 3}},
		{{to: 5, weight: 1}},
		{},
		{{to: 3, weight: 3}, {to: 7, weight: 1}},
		{{to: 3, weight: 5}, {to: 6, weight: 2}},
		{{to: 3, weight: 2}},
		{{to: 3, weight: 1}},
	}

	for _, tc := range testCases {
		got := g.ShortestPath(tc.from, tc.to)
		if got != tc.expected {
			t.Log(g.TopologicallySortedNodes())
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
