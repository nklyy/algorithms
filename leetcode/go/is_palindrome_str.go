package main

import (
	"strings"
	"unicode"
)

func isPalindromeStr(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		if !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l])) {
			l++
			continue
		} else if !unicode.IsLetter(rune(s[r])) && !unicode.IsDigit(rune(s[r])) {
			r--
			continue
		}

		// strings.ToLower(string(s[l])) == strings.ToLower(string(s[r])) -- more quick than strings.EqualFold(string(s[l]), string(s[r]))
		if strings.EqualFold(string(s[l]), string(s[r])) {
			l++
			r--
		} else {
			return false
		}

	}

	return true
}
