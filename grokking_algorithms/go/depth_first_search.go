package main

import "fmt"

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
			foundSeller := depthFirstSearch(graph, person)

			if foundSeller {
				return true
			}
		}
	}

	return false
}
