package main

import "math"

func sortedSquares(nums []int) []int {
	l := 0
	r := len(nums) - 1

	var res = make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		if math.Abs(float64(nums[l])) > math.Abs(float64(nums[r])) {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
	}

	return res
}
