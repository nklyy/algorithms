package main

func countElements(arr []int) int {
	checkMap := map[int]bool{}
	for i := 0; i < len(arr); i++ {
		checkMap[arr[i]] = true
	}

	ans := 0
	for i := 0; i < len(arr); i++ {
		if _, ok := checkMap[arr[i]+1]; ok {
			ans++
		}
	}

	return ans
}
