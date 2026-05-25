package main

func minSubArrayLen(target int, nums []int) int {
	l, ans, curr := 0, 0, 0

	for r := 0; r < len(nums); r++ {
		curr += nums[r]

		for curr >= target {
			if ans == 0 {
				ans = r - l + 1
			} else {
				ans = min(ans, r-l+1)
			}

			curr -= nums[l]
			l++
		}
	}

	return ans
}

// Alternative solution
// func minSubArrayLen(target int, nums []int) int {
// 	l, curr, ans := 0, 0, math.MaxInt32

// 	for right := 0; right < len(nums); right++ {
// 		curr += nums[right]

// 		for curr >= target {
// 			ans = min(ans, right-l+1)
// 			curr -= nums[l]
// 			l++
// 		}
// 	}

// 	if ans == math.MaxInt32 {
// 		return 0
// 	}
// 	return ans
// }
