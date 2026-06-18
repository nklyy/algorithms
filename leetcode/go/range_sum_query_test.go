package main

import "testing"

func TestRangeSumQuery(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	if got := obj.SumRange(0, 2); got != 1 {
		t.Errorf("SumRange(0, 2) = %d; want 1", got)
	}
	if got := obj.SumRange(2, 5); got != -1 {
		t.Errorf("SumRange(2, 5) = %d; want -1", got)
	}
	if got := obj.SumRange(0, 5); got != -3 {
		t.Errorf("SumRange(0, 5) = %d; want -3", got)
	}
}
