package main

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		s        []byte
		expected []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("Hannah"), []byte("hannaH")},
	}
	for _, test := range tests {
		reverseString(test.s)
		if string(test.s) != string(test.expected) {
			t.Errorf("ReverseString(%q) = %q; expected %q", test.s, test.s, test.expected)
		}
	}
}
