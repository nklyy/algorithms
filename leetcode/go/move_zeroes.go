package main

func moveZeroes(nums []int) {
	l := 0
	r := len(nums) - 1

	for l < r {
		if nums[l] == 0 {
			copy(nums, append(nums[:l], append(nums[l+1:], nums[l:l+1]...)...))
			r--
		} else {
			l++
		}
	}
}
