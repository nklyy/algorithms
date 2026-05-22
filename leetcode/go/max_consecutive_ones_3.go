package main

func longestOnes(nums []int, k int) int {
	l, z, ans := 0, 0, 0

	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			z++
		}

		for z > k {
			if nums[l] == 0 {
				z--
			}
			l++
		}

		ans = max(ans, r-l+1)
	}

	return ans
}
