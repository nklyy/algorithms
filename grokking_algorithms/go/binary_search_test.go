package main

import "testing"

func Test_binarySearchInts(t *testing.T) {
	type args struct {
		list []int64
		i    int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search FOUND",
			args: args{
				list: []int64{2, 4, 5},
				i:    2,
			},
			want: 0,
		},
		{
			name: "search NOT FOUND",
			args: args{
				list: []int64{2, 4, 5},
				i:    6,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.list, tt.args.i); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchStrings(t *testing.T) {
	type args struct {
		list []string
		i    string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search FOUND",
			args: args{
				list: []string{"a", "b", "c"},
				i:    "b",
			},
			want: 1,
		},
		{
			name: "search FOUND",
			args: args{
				list: []string{"a", "b", "c"},
				i:    "d",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.list, tt.args.i); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
