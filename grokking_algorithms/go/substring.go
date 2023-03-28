package main

import (
	"fmt"
)

func substring(a, b string) [][]int {
	cell := createMatrix(len(a), len(b))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				if i > 0 && j > 0 {
					cell[i][j] = cell[i-1][j-1] + 1
				} else {
					cell[i][j] = 0
				}
			} else {
				cell[i][j] = 0
			}
		}
	}

	for _, i := range cell {
		fmt.Println(i)
	}
	return cell
}
