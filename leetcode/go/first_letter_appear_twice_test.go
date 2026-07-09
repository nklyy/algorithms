package main

import "testing"

func TestRepeatedCharacter(t *testing.T) {
	tests := []struct {
		s        string
		expected byte
	}{
		{"abccbaacz", 'c'},
		{"abcdd", 'd'},

		// Test for a string with no repeated characters
		{"abcdefg", byte(0)},

		// Test for a string with all characters the same
		{"aaaaaa", 'a'},

		// Test for a string with special characters
		{"!@#$%^&*()!@#", '!'},

		// Test for a string with numbers
		{"12345678901234567890", '1'},

		// Test for a string with mixed characters
		{"a1b2c3d4e5f6g7h8i9j0a", 'a'},

		// Test for an empty string
		{"", byte(0)},
	}

	for _, test := range tests {
		result := repeatedCharacter(test.s)
		if result != test.expected {
			t.Errorf("For input %s, expected %c but got %c", test.s, test.expected, result)
		}
	}
}
