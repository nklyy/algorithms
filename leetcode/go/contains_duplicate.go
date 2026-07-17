package main

func containsDuplicate(nums []int) bool {
	checkMap := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := checkMap[nums[i]]; ok {
			return true
		}

		checkMap[nums[i]] = true
	}

	return false
}
