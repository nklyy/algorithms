package main

import (
	"reflect"
	"testing"
)

func Test_selectionSortInt(t *testing.T) {
	type args struct {
		lst []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				lst: []int{12, 2, 77, 33, 1, 0},
			},
			want: []int{0, 1, 2, 12, 33, 77},
		},
		{
			name: "test 2",
			args: args{
				lst: []int{-12, -2, -77, -33, -1},
			},
			want: []int{-77, -33, -12, -2, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectionSort(tt.args.lst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})

	}
}

func Test_selectionSortFloat(t *testing.T) {
	type args struct {
		lst []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "test 1",
			args: args{
				lst: []float32{0.2, 0.1, 0.5},
			},
			want: []float32{0.1, 0.2, 0.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectionSort(tt.args.lst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})

	}
}

func Test_selectionSortString(t *testing.T) {
	type args struct {
		lst []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test 1",
			args: args{
				lst: []string{"c", "b", "f", "a"},
			},
			want: []string{"a", "b", "c", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectionSort(tt.args.lst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})

	}
}
