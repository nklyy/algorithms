package main

import (
	"testing"
)

func TestEqualSubstring(t *testing.T) {
	tests := []struct {
		s       string
		t       string
		maxCost int
		want    int
	}{
		{"abcd", "bcdf", 3, 3},
		{"abcd", "cdef", 3, 1},
		{"abcd", "acde", 0, 1},
	}

	for _, tt := range tests {
		if got := equalSubstring(tt.s, tt.t, tt.maxCost); got != tt.want {
			t.Errorf("equalSubstring(%v, %v, %v) = %v; want %v", tt.s, tt.t, tt.maxCost, got, tt.want)
		}
	}
}
