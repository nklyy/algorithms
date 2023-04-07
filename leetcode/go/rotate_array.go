package main

func rotate(nums []int, k int) {
	// n := len(nums)
	// // k %= n

	// for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
	// 	nums[i], nums[j] = nums[j], nums[i]
	// }

	// for i, j := 0, (k - 1); i < j; i, j = i+1, j-1 {
	// 	nums[i], nums[j] = nums[j], nums[i]
	// }

	// for i, j := k, n-1; i < j; i, j = i+1, j-1 {
	// 	nums[i], nums[j] = nums[j], nums[i]
	// }

	copy(nums, append(nums[len(nums)-(k%len(nums)):], nums[:len(nums)-(k%len(nums))]...))
}
