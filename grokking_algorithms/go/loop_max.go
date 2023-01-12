package main

func loopMax(lst []int) int {
	max := lst[0]

	for i := 1; i < len(lst); i++ {
		if lst[i] > max {
			max = lst[i]
		}
	}

	return max
}
