package main

func findMiddleIndex(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	left := 0
	for i, val := range nums {
		if left == total-left-val {
			return i
		}
		left += val
	}

	return -1
}
