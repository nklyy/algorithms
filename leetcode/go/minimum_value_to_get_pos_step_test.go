package main

import (
	"reflect"
	"testing"
)

func TestMinStartValue(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test1",
			nums: []int{-3, 2, -3, 4, 2},
			want: 5,
		},
		{
			name: "test2",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "test3",
			nums: []int{1, -2, -3},
			want: 5,
		},
		{
			name: "test4",
			nums: []int{2, 3, 5, -5, -1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStartValue(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minStartValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
