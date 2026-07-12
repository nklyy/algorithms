package main

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Input: %v", tt.nums), func(t *testing.T) {
			if got := missingNumber(tt.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
