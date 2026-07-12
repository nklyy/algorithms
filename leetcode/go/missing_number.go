package main

func missingNumber(nums []int) int {
	checkMap := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		checkMap[nums[i]] = true
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := checkMap[i]; !ok {
			return i
		}
	}

	return 0
}
