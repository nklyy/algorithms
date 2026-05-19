package main

import "testing"

func TestGetCommon(t *testing.T) {
	tests := []struct {
		n1       []int
		n2       []int
		expected int
	}{
		{n1: []int{1, 2, 3}, n2: []int{2, 3}, expected: 2},
		{n1: []int{1, 2, 3}, n2: []int{4, 5, 6}, expected: -1},
		{n1: []int{1, 2}, n2: []int{2, 4}, expected: 2},
	}
	for _, test := range tests {
		if got := getCommon(test.n1, test.n2); got != test.expected {
			t.Errorf("getCommon(%v, %v) = %v, want %v", test.n1, test.n2, got, test.expected)
		}
	}
}
