package main

import "testing"

func TestReversePrefix(t *testing.T) {
	tests := []struct {
		s        string
		ch       byte
		expected string
	}{
		{"abcdefd", 'd', "dcbaefd"},
		{"xyxzxe", 'z', "zxyxxe"},
		{"abcd", 'z', "abcd"},
	}
	for _, test := range tests {
		if got := reversePrefix(test.s, test.ch); got != test.expected {
			t.Errorf("reversePrefix(%v, %v) = %v, want %v", test.s, test.ch, got, test.expected)
		}
	}
}
