package main

import "fmt"

func destCity(paths [][]string) string {
	mapPath := map[string]string{}
	for i := 0; i < len(paths); i++ {
		mapPath[paths[i][0]] = paths[i][1]
	}

	fmt.Println(mapPath)
	for i := 0; i < len(paths); i++ {
		if _, ok := mapPath[paths[i][1]]; !ok {
			return paths[i][1]
		}
	}

	return ""
}
