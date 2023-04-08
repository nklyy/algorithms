package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		lst []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				lst: []int{1, 2, 3, 4, 5, 6, 7},
				n:   3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "2",
			args: args{
				lst: []int{-1, -100, 3, 99},
				n:   2,
			},
			want: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.lst, tt.args.n)

			if !reflect.DeepEqual(tt.args.lst, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.args.lst, tt.want)
			}
		})
	}
}
