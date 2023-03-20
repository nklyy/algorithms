package main

import (
	"fmt"
)

func isSellerDfs(name string) bool {
	if name == "you" {
		return false
	}

	return name[len(name)-1] == 'm'
}

func isSearchedDfs(person string, searched []string) bool {
	for _, n := range searched {
		if n == person {
			return true
		}
	}
	return false
}

var searched []string

func depthFirstSearch(graph map[string][]string, name string) bool {
	if isSellerDfs(name) {
		fmt.Println("You fund mango seller:", name)
		return true
	}

	if isSearchedDfs(name, searched) {
		return false
	}

	searched = append(searched, name)

	for _, person := range graph[name] {
		if !isSearchedDfs(person, searched) {
			if foundSeller := depthFirstSearch(graph, person); foundSeller {
				return true
			}
		}
	}

	return false
}

func depthFirstSearchV2(graph map[string][]string, v, t string) bool {
	if v == t {
		return true
	}

	if isSearchedDfs(v, searched) {
		return false
	}

	searched = append(searched, v)

	for _, neighbor := range graph[v] {
		if !isSearchedDfs(neighbor, searched) {
			if foundSeller := depthFirstSearchV2(graph, neighbor, t); foundSeller {
				return true
			}
		}
	}

	return false
}
