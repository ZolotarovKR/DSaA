package shortest_path

import (
	"math"
	"slices"
)

var graph = [][]struct {
	to     int
	weight float64
}{
	{{4, 2}},
	{{0, 1}, {2, 1}, {4, 4}, {5, 3}},
	{{5, 1}},
	{},
	{{3, 3}, {7, 1}},
	{{3, 5}, {6, 2}},
	{{3, 2}},
	{{3, 1}},
}

func dfs(from int, nodes *[]int, visited []bool) {
	visited[from] = true
	for _, edge := range graph[from] {
		if !visited[edge.to] {
			dfs(edge.to, nodes, visited)
		}
	}
	*nodes = append(*nodes, from)
}

func topologicalSort() []int {
	nodes := make([]int, 0, len(graph))
	visited := make([]bool, len(graph))
	for i := range graph {
		if !visited[i] {
			dfs(i, &nodes, visited)
		}
	}
	slices.Reverse(nodes)
	return nodes
}

func ShortestPath(from, to int) float64 {
	ds := make([]float64, len(graph))
	for i := range ds {
		ds[i] = math.Inf(1)
	}
	ds[from] = 0
	nodes := topologicalSort()
	for _, n := range nodes {
		for _, e := range graph[n] {
			if ds[n]+e.weight < ds[e.to] {
				ds[e.to] = ds[n] + e.weight
			}
		}
	}
	return ds[to]
}
