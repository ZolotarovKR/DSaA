package warshall

func Warshall(graph [][]bool) [][]bool {
	t := make([][]bool, len(graph))
	for j, r := range graph {
		t[j] = make([]bool, len(r))
		copy(t[j], r)
	}
	for k := 0; k < len(graph); k++ {
		for j, r := range t {
			for i, v := range r {
				t[j][i] = v || (t[j][k] && t[k][i])
			}
		}
	}
	return t
}
