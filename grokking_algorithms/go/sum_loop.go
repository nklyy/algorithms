package main

func sumLoop(arr []int) int {
	r := 0

	for i := 0; i < len(arr); i++ {
		r += arr[i]
	}

	return r
}
