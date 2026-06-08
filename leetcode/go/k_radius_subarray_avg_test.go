package main

import (
	"reflect"
	"testing"
)

func TestGetAverages(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "test1",
			nums: []int{7, 4, 3, 9, 1, 8, 5, 2, 6},
			k:    3,
			want: []int{-1, -1, -1, 5, 4, 4, -1, -1, -1},
		},
		{
			name: "test2",
			nums: []int{100000},
			k:    0,
			want: []int{100000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAverages(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
