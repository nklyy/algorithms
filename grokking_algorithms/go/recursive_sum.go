package main

func recursiveSum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + recursiveSum(arr[1:])
}
