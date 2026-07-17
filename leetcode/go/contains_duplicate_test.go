package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}, true},
	}

	for _, tc := range testCases {
		result := containsDuplicate(tc.nums)
		if result != tc.expected {
			t.Errorf("containsDuplicate(%v) = %v; expected %v", tc.nums, result, tc.expected)
		}
	}
}
