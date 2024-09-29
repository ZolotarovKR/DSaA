package shortest_path

import (
	"math"
	"slices"
)

type Graph [][]struct {
	to     int
	weight float64
}

func (g Graph) TopologicallySortedNodes() []int {
	nodes := make([]int, 0, len(g))
	visited := make([]bool, len(g))
	var dfs func(int)
	dfs = func(from int) {
		visited[from] = true
		for _, e := range g[from] {
			if !visited[e.to] {
				dfs(e.to)
			}
		}
		nodes = append(nodes, from)
	}
	for i := range g {
		if !visited[i] {
			dfs(i)
		}
	}
	slices.Reverse(nodes)
	return nodes
}

func (g Graph) ShortestPath(from, to int) float64 {
	ds := make([]float64, len(g))
	for i := range ds {
		ds[i] = math.Inf(1)
	}
	ds[from] = 0
	nodes := g.TopologicallySortedNodes()
	for _, n := range nodes {
		for _, e := range g[n] {
			if ds[n]+e.weight < ds[e.to] {
				ds[e.to] = ds[n] + e.weight
			}
		}
	}
	return ds[to]
}
