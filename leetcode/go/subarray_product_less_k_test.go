package main

import "testing"

func TestNumSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{10, 5, 2, 6}, 100, 8},
		{[]int{1, 2, 3}, 0, 0},
		{[]int{1, 2, 3}, 1, 0},
		{[]int{1, 2, 3}, 2, 1},
	}

	for _, test := range tests {
		if got := numSubarrayProductLessThanK(test.nums, test.k); got != test.expected {
			t.Errorf("numSubarrayProductLessThanK(%v, %d) = %d; want %d", test.nums, test.k, got, test.expected)
		}
	}
}
