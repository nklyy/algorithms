package main

func runningSum(nums []int) []int {
	ans := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		ans = append(ans, ans[len(ans)-1]+nums[i])
	}

	return ans
}
