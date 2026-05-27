package main

import (
	"math"
)

func equalSubstring(s string, t string, maxCost int) int {
	b1 := []byte(s)
	b2 := []byte(t)

	l, curr, ans := 0, 0, 0

	for r := 0; r < len(b1); r++ {
		curr += int(math.Abs(float64(int(b1[r]) - int(b2[r]))))

		for curr > maxCost {
			curr -= int(math.Abs(float64(int(b1[l]) - int(b2[l]))))
			l++
		}

		ans = max(ans, r-l+1)
	}

	return ans
}
