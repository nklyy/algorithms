package main

import "fmt"

func breadthFirstSearch(graph map[string][]string, name string) bool {
	// var queue []string
	// queue = append(queue, graph[name]...)

	queue := graph[name]

	var visited []string

	for len(queue) != 0 {
		person := queue[0]
		queue = queue[1:]

		if !contains(person, visited) {
			if isSeller(person) {
				fmt.Println("You fund mango seller:", person)
				return true
			}

			queue = append(queue, graph[person]...)
			visited = append(visited, person)
		}
	}

	return false
}

func breadthFirstSearchV2[K comparable](graph map[K][]K, v, t K) bool {
	queue := []K{v}
	visited := []K{}

	for len(queue) != 0 {
		neighbor := queue[0]
		queue = queue[1:]

		if !contains(neighbor, visited) {
			if neighbor == t {
				return true
			}

			queue = append(queue, graph[neighbor]...)
			visited = append(visited, neighbor)
		}
	}

	return false
}
