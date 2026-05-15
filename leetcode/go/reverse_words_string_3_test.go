package main

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"Hello World", "olleH dlroW"},
		{"a", "a"},
	}
	for _, test := range tests {
		if got := reverseWords(test.s); got != test.expected {
			t.Errorf("reverseWords(%q) = %q; expected %q", test.s, got, test.expected)
		}
	}
}
