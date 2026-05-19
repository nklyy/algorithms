package main

// func moveZeroes(nums []int) {
// 	l := 0
// 	r := len(nums) - 1

// 	for l < r {
// 		if nums[l] == 0 {
// 			copy(nums, append(nums[:l], append(nums[l+1:], nums[l:l+1]...)...))
// 			r--
// 		} else {
// 			l++
// 		}
// 	}
// }

// func moveZeroes(nums []int) {
// 	if len(nums) == 1 {
// 		return
// 	}

// 	for i := 0; i <= len(nums); i++ {
// 		l := 0
// 		r := 1

// 		for r < len(nums) {
// 			if nums[l] == 0 && nums[r] != 0 {
// 				nums[l], nums[r] = nums[r], nums[l]
// 				l++
// 				r++
// 				continue
// 			}

// 			l++
// 			r++
// 		}
// 	}
// }

func moveZeroes(nums []int) {
	h := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[h], nums[i] = nums[i], nums[h]
			h++
		}
	}
}
