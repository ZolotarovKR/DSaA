package shortest_safe_path

import (
	"math"
	"slices"
)

type Graph [][]struct {
	to     int
	weight float64
}

func ShortestSafePath(g Graph, from, to, lim int) float64 {
	d := make([][]float64, len(g))
	for j := range d {
		d[j] = make([]float64, lim+1)
		for i := range d[j] {
			d[j][i] = math.Inf(1)
		}
	}
	d[from][0] = 0
	for k := 1; k <= lim; k++ {
		for v, r := range g {
			for _, e := range r {
				d[e.to][k] = math.Min(d[e.to][k], d[v][k-1]+e.weight)
			}
		}
	}
	return slices.Min(d[to])
}
