package main

import (
	"fmt"
)

func depthFirstSearch(graph map[string][]string, name string, visited []string) bool {
	if isSeller(name) {
		fmt.Println("You fund mango seller:", name)
		return true
	}

	if contains(name, visited) {
		return false
	}

	visited = append(visited, name)

	for _, neighbor := range graph[name] {
		if !contains(neighbor, visited) {
			if foundSeller := depthFirstSearch(graph, neighbor, visited); foundSeller {
				return true
			}
		}
	}

	return false
}

func depthFirstSearchV2[K comparable](graph map[K][]K, v, t K, visited []K) bool {
	if v == t {
		return true
	}

	if contains(v, visited) {
		return false
	}

	visited = append(visited, v)

	for _, neighbor := range graph[v] {
		if !contains(neighbor, visited) {
			if found := depthFirstSearchV2(graph, neighbor, t, visited); found {
				return true
			}
		}
	}

	return false
}
