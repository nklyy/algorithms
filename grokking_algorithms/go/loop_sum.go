package main

func loopSum(lst []int) int {
	r := 0

	for _, v := range lst {
		r += v
	}

	return r
}
