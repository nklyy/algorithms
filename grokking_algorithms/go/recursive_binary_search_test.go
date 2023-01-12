package main

import "testing"

func Test_recursiveBinarySearchRecursiveInts(t *testing.T) {
	type args struct {
		lst []int
		i   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "found",
			args: args{
				lst: []int{2, 4, 5},
				i:   2,
			},
			want: 0,
		},
		{
			name: "not found",
			args: args{
				lst: []int{2, 4, 5},
				i:   6,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursiveBinarySearchRecursive(tt.args.lst, tt.args.i, 0, len(tt.args.lst)-1); got != tt.want {
				t.Errorf("recursiveBinarySearchRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recursiveBinarySearchRecursiveStrings(t *testing.T) {
	type args struct {
		lst []string
		i   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "found",
			args: args{
				lst: []string{"a", "b", "c"},
				i:   "b",
			},
			want: 1,
		},
		{
			name: "not found",
			args: args{
				lst: []string{"a", "b", "c"},
				i:   "d",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursiveBinarySearchRecursive(tt.args.lst, tt.args.i, 0, len(tt.args.lst)-1); got != tt.want {
				t.Errorf("recursiveBinarySearchRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
