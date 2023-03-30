package main

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		lst []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				lst: []int{1, 3, 5, 6},
				n:   5,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				lst: []int{1, 3, 5, 6},
				n:   7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.lst, tt.args.n); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
