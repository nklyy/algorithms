package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	res := 0
	n := x

	for n != 0 {
		remain := n % 10
		res = (res * 10) + remain

		n /= 10
	}

	return x == res
}
