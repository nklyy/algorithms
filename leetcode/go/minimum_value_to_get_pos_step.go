package main

func minStartValue(nums []int) int {
	prefix := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		prefix = append(prefix, prefix[len(prefix)-1]+nums[i])
	}

	min := prefix[0]
	for i := 1; i < len(prefix); i++ {
		if prefix[i] < min {
			min = prefix[i]
		}
	}

	if min > 0 {
		return 1
	}

	return -min + 1
}

// optimized version
// func minStartValue(nums []int) int {
// 	total := 0
// 	minVal := 0
// 	for i := 0; i < len(nums); i++ {
// 		total += nums[i]
// 		minVal = min(minVal, total)
// 	}

// 	return -minVal + 1
// }
