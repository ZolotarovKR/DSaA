package lcs

import "testing"

func TestLcs(t *testing.T) {
	testCases := []struct {
		xs, ys   []rune
		expected int
	}{
		{
			[]rune("ACGGAGTGCGCGGAAGCCGGGCCGAA"),
			[]rune("GTTCGGAATGCCGTTGCTCTGTAA"),
			17,
		},
		{
			[]rune("BDCABA"),
			[]rune("ABCBDAB"),
			4,
		},
	}

	for _, tc := range testCases {
		got := Lcs(tc.xs, tc.ys)
		if len(got) != tc.expected {
			t.Logf("xs: %v", string(tc.xs))
			t.Logf("ys: %v", string(tc.ys))
			t.Logf("got: %v (length: %v)", string(got), len(got))
			t.Fatalf(
				"Expected length of LCS is %v, but actual is %v",
				tc.expected,
				len(got),
			)
		}

	}
}
