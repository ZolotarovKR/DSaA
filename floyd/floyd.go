package floyd

func Floyd(graph [][]float64) [][]float64 {
	d := make([][]float64, len(graph))
	for j, r := range graph {
		d[j] = make([]float64, len(r))
		copy(d[j], r)
	}
	for k := 0; k < len(graph); k++ {
		for j, r := range d {
			for i, v := range r {
				t := d[j][k] + d[k][i]
				if t < v {
					d[j][i] = t
				}
			}
		}
	}
	return d
}
