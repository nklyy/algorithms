package main

import "testing"

func TestIsPalindromeStr(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
	}
	for _, test := range tests {
		actual := isPalindromeStr(test.s)
		if actual != test.expected {
			t.Errorf("IsPalindromeStr(%q) = %v; expected %v", test.s, actual, test.expected)
		}
	}
}
