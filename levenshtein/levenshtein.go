package levenshtein

func Levenshtein(s1, s2 string) int {
	xs, ys := []rune(s1), []rune(s2)
	m := make([][]int, len(xs)+1)
	for j := range m {
		m[j] = make([]int, len(ys)+1)
	}
	for j := 1; j <= len(xs); j++ {
		m[j][0] = j
	}
	for i := 1; i <= len(ys); i++ {
		m[0][i] = i
	}
	for j := 1; j <= len(xs); j++ {
		for i := 1; i <= len(ys); i++ {
			if xs[j-1] == ys[i-1] {
				m[j][i] = m[j-1][i-1]
			} else {
				m[j][i] = m[j-1][i] + 1
				if m[j-1][i-1]+1 < m[j][i] {
					m[j][i] = m[j-1][i-1] + 1
				}
				if m[j][i-1]+1 < m[j][i] {
					m[j][i] = m[j][i-1] + 1
				}
			}
		}
	}
	return m[len(xs)][len(ys)]
}
