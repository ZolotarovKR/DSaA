package lcs

import "slices"

func Lcs(xs, ys []rune) []rune {
	m := make([][]int, len(xs)+1)
	for j := range m {
		m[j] = make([]int, len(ys)+1)
	}
	for j := 1; j < len(m); j++ {
		for i := 1; i < len(m[j]); i++ {
			if xs[j-1] == ys[i-1] {
				m[j][i] = m[j-1][i-1] + 1
			} else {
				if m[j-1][i] >= m[j][i-1] {
					m[j][i] = m[j-1][i]
				} else {
					m[j][i] = m[j][i-1]
				}
			}
		}
	}
	rs := make([]rune, m[len(xs)][len(ys)])
	i, j := len(ys), len(xs)
	k := 0
	for k < len(rs) {
		if xs[j-1] == ys[i-1] {
			rs[k] = xs[j-1]
			j -= 1
			i -= 1
			k += 1
		} else {
			if m[j][i] == m[j-1][i] {
				j -= 1
			} else {
				i -= 1
			}
		}
	}
	slices.Reverse(rs)
	return rs
}
