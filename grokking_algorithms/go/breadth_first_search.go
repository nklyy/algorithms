package main

import "fmt"

func isSeller(name string) bool {
	return name[len(name)-1] == 'm'
}

func isSearched(person string, searched []string) bool {
	for _, n := range searched {
		if n == person {
			return true
		}
	}
	return false
}

func breadthSearch(graph map[string][]string, name string) bool {
	// var queue []string
	// queue = append(queue, graph[name]...)

	queue := graph[name]

	var searched []string

	for len(queue) != 0 {
		person := queue[0]
		queue = queue[1:]

		if !isSearched(person, searched) {
			if isSeller(person) {
				fmt.Println("You fund mango seller:", person)
				return true
			}

			queue = append(queue, graph[person]...)
			searched = append(searched, person)
		}
	}

	return false
}
