package main

type Graph struct {
	graph map[int][]int
	V     int // number of vertices
}

func NewGraph(v int) Graph {
	return Graph{
		graph: make(map[int][]int),
		V:     v,
	}
}

func (g *Graph) addEdge(u, v int) {
	g.graph[u] = append(g.graph[u], v)
}

func (g *Graph) topologySortUtil(v int, visited []bool, stack *[]int) {
	visited[v] = true

	for _, i := range g.graph[v] {
		if !visited[i] {
			g.topologySortUtil(i, visited, stack)
		}
	}

	*stack = append(*stack, v)
}

func (g *Graph) topologySort() []int {
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
