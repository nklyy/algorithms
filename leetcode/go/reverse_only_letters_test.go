package main

import "testing"

func TestReverseOnlyLetters(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"ab-cd", "dc-ba"},
		{"a-bC-dEf-ghIj", "j-Ih-gfE-dCba"},
		{"Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!"},
	}
	for _, test := range tests {
		if got := reverseOnlyLetters(test.s); got != test.expected {
			t.Errorf("ReverseOnlyLetters(%q) = %q; expected %q", test.s, got, test.expected)
		}
	}
}
