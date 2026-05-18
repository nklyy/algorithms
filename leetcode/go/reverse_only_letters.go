package main

import "unicode"

func reverseOnlyLetters(s string) string {
	rn := []rune(s)
	l := 0
	r := len(rn) - 1

	for l < r {
		if !unicode.IsLetter(rn[l]) {
			l++
			continue
		} else if !unicode.IsLetter(rn[r]) {
			r--
			continue
		}

		rn[l], rn[r] = rn[r], rn[l]
		l++
		r--
	}

	return string(rn)
}
