package main

// func pivotIndex(nums []int) int {
// 	p := []int{nums[0]}
// 	for i := 1; i < len(nums); i++ {
// 		p = append(p, p[len(p)-1]+nums[i])
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		if i == 0 {
// 			if p[len(p)-1]-p[i] == 0 {
// 				return i
// 			}
// 			continue
// 		}

// 		if i == len(nums)-1 {
// 			if p[i-1] == 0 {
// 				return i
// 			}
// 			continue
// 		}

// 		if p[i-1] == p[len(p)-1]-p[i] {
// 			return i
// 		}
// 	}

// 	return -1
// }

// Alternative solution using a single pass
func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	leftSum := 0
	for i, num := range nums {
		if leftSum == total-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}
