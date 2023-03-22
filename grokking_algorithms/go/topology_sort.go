package main

type TopologyGraph struct {
	graph map[int][]int
	V     int // number of vertices
}

func NewGraph(v int) TopologyGraph {
	return TopologyGraph{
		graph: make(map[int][]int),
		V:     v,
	}
}

func (g *TopologyGraph) addEdge(u, v int) {
	g.graph[u] = append(g.graph[u], v)
}

func (g *TopologyGraph) topologySortUtil(v int, visited []bool, stack *[]int) {
	visited[v] = true

	for _, i := range g.graph[v] {
		if !visited[i] {
			g.topologySortUtil(i, visited, stack)
		}
	}

	*stack = append(*stack, v)
}

func (g *TopologyGraph) topologySort() []int {
	visited := make([]bool, g.V)
	stack := []int{}

	for i := 0; i < g.V; i++ {
		if !visited[i] {
			g.topologySortUtil(i, visited, &stack)
		}
	}

	// reverse(stack)
	return stack
}

// func reverse[S ~[]E, E any](s S) {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }
