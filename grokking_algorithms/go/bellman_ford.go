package main

import (
	"fmt"
	"math"
)

type BellmanFordEdge struct {
	src   int
	dest  int
	weigh int
}

type BellmanFordGraph struct {
	graph []BellmanFordEdge
	V     int
}

func NewBellmanFordGraph(v int) BellmanFordGraph {
	return BellmanFordGraph{
		graph: make([]BellmanFordEdge, v),
		V:     v,
	}
}

func (g *BellmanFordGraph) addEdge(s, d, w int) {
	g.graph = append(g.graph, BellmanFordEdge{src: s, dest: d, weigh: w})
}

func (g *BellmanFordGraph) bellmanFord(src int) []int {
	var distances = make([]int, g.V)

	for i := 0; i < g.V; i++ {
		distances[i] = math.MaxInt64
	}

	distances[src] = 0

	for i := 0; i < (g.V - 1); i++ {
		for _, e := range g.graph {
			if distances[e.src] != math.MaxInt64 && distances[e.dest] > distances[e.src]+e.weigh {
				distances[e.dest] = distances[e.src] + e.weigh
			}
		}
	}

	for _, e := range g.graph {
		if distances[e.src] != math.MaxInt64 && distances[e.dest] > distances[e.src]+e.weigh {
			fmt.Println("Graph contains negative weight cycle")
			return []int{}
		}
	}

	// for i := 0; i < g.V; i++ {
	// 	fmt.Printf("%v - %v\n", i, distances[i])
	// }

	return distances
}
