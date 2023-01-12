package main

func recursiveSum(lst []int) int {
	if len(lst) == 1 {
		return lst[0]
	}

	return lst[0] + recursiveSum(lst[1:])
}
