package shortest_safe_path

import "testing"

func TestShortestSafePath(t *testing.T) {
	testCases := []struct {
		from, to int
		expected float64
	}{
		{1, 3, 7},
		{1, 7, 5},
		{0, 7, 3},
		{1, 6, 5},
	}
	lim := 2
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
		got := ShortestSafePath(g, tc.from, tc.to, lim)
		if got != tc.expected {
			t.Fatalf(
				"Expected ShortestSafePath(%v, %v) to equal %v, but got %v",
				tc.from,
				tc.to,
				tc.expected,
				got,
			)
		}
	}
}
