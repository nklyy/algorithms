package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	l, ans := 0, 0

	curr := 1

	for r := 0; r < len(nums); r++ {
		curr *= nums[r]

		for curr >= k {
			curr /= nums[l]
			l++
		}

		ans += r - l + 1
	}

	return ans
}
