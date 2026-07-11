package main

import "testing"

func TestCheckIfPangram(t *testing.T) {
	tests := []struct {
		sentence string
		want     bool
	}{
		{"thequickbrownfoxjumpsoverthelazydog", true},
		{"leetcode", false},
		{"thequickbrownfoxjumpsoverthelazydog", true},
	}

	for _, tt := range tests {
		t.Run(tt.sentence, func(t *testing.T) {
			if got := checkIfPangram(tt.sentence); got != tt.want {
				t.Errorf("checkIfPangram() = %v, want %v", got, tt.want)
			}
		})
	}
}
