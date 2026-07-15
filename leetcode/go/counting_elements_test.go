package main

import "testing"

// 1426 leetcode/go/counting_elements.go
func TestCountElements(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "test case 1",
			arr:  []int{1, 2, 3},
			want: 2,
		},
		{
			name: "test case 2",
			arr:  []int{1, 1, 3, 3, 5, 5, 7, 7},
			want: 0,
		},
		{
			name: "test case 3",
			arr:  []int{1, 3, 2, 3, 5, 0},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countElements(tt.arr); got != tt.want {
				t.Errorf("countElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
