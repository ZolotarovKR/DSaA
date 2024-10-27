package splitting

import "testing"

func TestSplit(t *testing.T) {
	testCases := []struct {
		weights           []int
		buckets, expected int
	}{
		{
			[]int{100, 200, 300, 400, 500, 600, 700, 800, 900},
			3,
			1700,
		},
	}

	for _, tc := range testCases {
		actual := Split(tc.weights, tc.buckets)
		if actual != tc.expected {
			t.Logf("weights: %v", tc.weights)
			t.Logf("buckets: %v", tc.buckets)
			t.Fatalf(
				"Expected weight of the heaviest bucket is %v, but actual is %v",
				tc.expected,
				actual,
			)
		}
	}
}
