package main

import (
	"math"
)

type QPair struct {
	node int
	w    int
}

type DijkstraEdge struct {
	to int
	w  int
}

type DijkstraGraph struct {
	graph [][]DijkstraEdge
	V     int // number of vertices
}

func NewDijkstraGraph(v int) DijkstraGraph {
	return DijkstraGraph{
		graph: make([][]DijkstraEdge, v),
		V:     v,
	}
}

func (g *DijkstraGraph) addEdge(s, d, w int) {
	g.graph[s] = append(g.graph[s], DijkstraEdge{to: d, w: w})

	// comment this line if you want to use (Directed graph)
	g.graph[d] = append(g.graph[d], DijkstraEdge{to: s, w: w})
}

func (g *DijkstraGraph) dijkstra(src int) []int {
	var distances = make([]int, g.V)

	for i := 0; i < g.V; i++ {
		distances[i] = math.MaxInt64
	}

	queue := []QPair{}
	queue = append(queue, QPair{node: src, w: 0})

	distances[src] = 0

	for len(queue) > 0 {
		n := queue[0].node
		queue = queue[1:]

		for _, edge := range g.graph[n] {
			if distances[edge.to] > distances[n]+edge.w {
				distances[edge.to] = distances[n] + edge.w
				queue = append(queue, QPair{node: edge.to, w: distances[edge.to]})
			}
		}
	}

	// for i := 0; i < g.V; i++ {
	// 	fmt.Printf("%v - %v\n", i, distances[i])
	// }

	// fmt.Println(distances[1:])
	return distances
}
