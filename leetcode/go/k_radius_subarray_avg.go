package main

import (
	"slices"
)

func getAverages(nums []int, k int) []int {
	prefix := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		prefix = append(prefix, prefix[len(prefix)-1]+nums[i])
	}

	window := k * 2
	ans := slices.Repeat([]int{-1}, len(nums))
	if window < len(nums) {
		for i := window; i < len(nums); i++ {
			ans[i-k] = (prefix[i] - prefix[i-window] + nums[i-window]) / (window + 1)
		}
	}

	return ans
}
