package main

func contains[K comparable](v K, arr []K) bool {
	for _, i := range arr {
		if i == v {
			return true
		}
	}
	return false
}

func isSeller(name string) bool {
	if name == "you" {
		return false
	}

	return name[len(name)-1] == 'm'
}

func createMatrix(rows, cols int) [][]int {
	cell := make([][]int, rows)
	for i := range cell {
		cell[i] = make([]int, cols)
	}

	return cell
}
