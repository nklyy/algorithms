package main

import (
	"reflect"
	"testing"
)

func Test_selectionSortInts(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				arr: []int{12, 2, 77, 33, 1, 0},
			},
			want: []int{0, 1, 2, 12, 33, 77},
		},
		{
			name: "test 2",
			args: args{
				arr: []int{-12, -2, -77, -33, -1},
			},
			want: []int{-77, -33, -12, -2, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})

	}
}

func Test_selectionSortStrings(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test 1",
			args: args{
				arr: []string{"c", "b", "f", "a"},
			},
			want: []string{"a", "b", "c", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})

	}
}
