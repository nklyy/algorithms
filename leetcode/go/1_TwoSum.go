package main

// O(n2) - nested loops
func twoSumV1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// O(n) - Two-pass Hash Table
func twoSumV2(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, v := range nums {
		hashmap[v] = i
	}

	for i, v := range nums {
		complement := target - v

		if _, ok := hashmap[complement]; ok && hashmap[complement] != i {
			return []int{i, hashmap[complement]}
		}
	}

	return nil
}

// O(n) - One-pass Hash Table
func twoSumV3(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, v := range nums {
		complement := target - v

		if _, ok := hashmap[complement]; ok {
			return []int{hashmap[complement], i}
		}

		hashmap[v] = i
	}

	return nil
}
