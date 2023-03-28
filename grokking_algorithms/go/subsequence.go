package main

import (
	"fmt"
)

func subsequence(a, b string) [][]int {
	cell := createMatrix(len(a), len(b))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				if i > 0 && j > 0 {
					cell[i][j] = cell[i-1][j-1] + 1
				} else {
					cell[i][j] = 1
				}
			} else {
				if i > 0 && j > 0 {
					cell[i][j] = Max(cell[i-1][j], cell[i][j-1])
				} else {
					cell[i][j] = 1
				}
			}
		}
	}

	for _, i := range cell {
		fmt.Println(i)
	}
	return cell
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
