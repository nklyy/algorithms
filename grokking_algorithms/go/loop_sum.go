package main

func loopSum(arr []int) int {
	r := 0

	for _, v := range arr {
		r += v
	}

	return r
}
