package main

// Original solution using sliding window technique time complexity O(n) and space complexity O(1)
func maxVowels(s string, k int) int {
	b := []byte(s)
	l, count, ans := 0, 0, 0

	for r := 0; r < len(b); r++ {
		if isVowel(b[r]) {
			count++
		}

		if r-l >= k {
			if isVowel(b[l]) {
				count--
			}

			l++
		}

		ans = max(ans, count)
	}

	return ans
}

// Alternative solution using sliding window technique without using max function time complexity O(n) and space complexity O(1)
// func maxVowels(s string, k int) int {
// 	b := []byte(s)
// 	curr := 0

// 	for i := 0; i < k; i++ {
// 		if isVowel(b[i]) {
// 			curr++
// 		}
// 	}

// 	ans := curr
// 	for i := k; i < len(b); i++ {
// 		if isVowel(b[i]) {
// 			curr++
// 		}

// 		if isVowel(b[i-k]) {
// 			curr--
// 		}

// 		if curr > ans {
// 			ans = curr
// 		}
// 	}

// 	return ans
// }

func isVowel(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
		return true
	}

	return false
}
