package main

import "testing"

func TestFindMiddleIndex(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, -1, 8, 4}, 3},
		{[]int{1, -1, 4}, 2},
		{[]int{2, 5}, -1},
	}
	for _, test := range tests {
		if got := findMiddleIndex(test.nums); got != test.expected {
			t.Errorf("findMiddleIndex(%v) = %v, want %v", test.nums, got, test.expected)
		}
	}
}
