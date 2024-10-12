package lis

import (
	"cmp"
	"slices"
)

func Lis[T cmp.Ordered](xs []T) []T {
	if len(xs) == 0 {
		return []T{}
	}
	ds := make([]int, len(xs))
	ps := make([]int, len(xs))
	for i := range xs {
		ds[i] = 1
		ps[i] = -1
	}
	for i := 1; i < len(ds); i++ {
		for j := 0; j < i; j++ {
			if xs[j] < xs[i] && ds[j]+1 > ds[i] {
				ds[i] = ds[j] + 1
				ps[i] = j
			}
		}
	}
	last_idx := 0
	for i := 1; i < len(ds); i++ {
		if ds[i] > ds[last_idx] {
			last_idx = i
		}
	}
	lis := make([]T, 0, len(xs)>>2)
	for i := last_idx; i != -1; i = ps[i] {
		lis = append(lis, xs[i])
	}
	slices.Reverse(lis)
	return lis
}
