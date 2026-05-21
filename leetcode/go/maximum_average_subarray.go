package main

func findMaxAverage(nums []int, k int) float64 {
	curr := 0

	for i := 0; i < k; i++ {
		curr += nums[i]
	}

	ans := curr
	for i := k; i < len(nums); i++ {
		curr += nums[i] - nums[i-k]

		ans = max(ans, curr)
	}

	return float64(ans) / float64(k)
}
