package main

import "testing"

func Test_recursiveMax(t *testing.T) {
	type args struct {
		lst []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				lst: []int{25, 4, 26},
			},
			want: 26,
		},
		{
			name: "test 2",
			args: args{
				lst: []int{-2, 4, -6},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursiveMax(tt.args.lst); got != tt.want {
				t.Errorf("recursiveMax() = %v, want %v", got, tt.want)
			}
		})

	}
}
