package main

func recursiveMax(arr []int) int {
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		}
		return arr[1]
	}

	max := recursiveMax(arr[1:])
	if arr[0] > max {
		return arr[0]
	}

	return max
}
