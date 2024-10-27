package splitting

import (
	"math"
	"slices"
)

func maxInt(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func minInt(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func Split(weights []int, buckets int) int {
	if len(weights) == buckets {
		return slices.Max(weights)
	}

	if buckets == 1 {
		sum := 0
		for _, w := range weights {
			sum += w
		}
		return sum
	}

	r := math.MaxInt
	for i := 1; i <= len(weights)-buckets+1; i++ {
		s := Split(weights[:len(weights)-i], buckets-1)
		t := 0
		for j := i; j < len(weights); j++ {
			t += weights[j]
		}
		r = minInt(r, maxInt(s, t))

	}
	return r
}
