package levenshtein

import "testing"

func TestLevenshtein(t *testing.T) {
	testCases := []struct {
		s1, s2   string
		expected int
	}{
		{"book", "back", 2},
		{"fiber", "fish", 3},
		{"horse", "house", 1},
		{"algorithm", "algorithm", 0},
	}

	for _, tc := range testCases {
		actual := Levenshtein(tc.s1, tc.s2)
		if actual != tc.expected {
			t.Logf("s1: %v", tc.s1)
			t.Logf("s2: %v", tc.s2)
			t.Fatalf(
				"Expected distance is %v, but actual is %v",
				tc.expected,
				actual,
			)
		}
	}
}
