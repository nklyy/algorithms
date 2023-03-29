package main

import "testing"

func Test_binarySearch(t *testing.T) {
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
			if got := search(tt.args.lst, tt.args.i); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
