package main

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		lst []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				lst: []int{-4, -1, 0, 3, 10},
			},
			want: []int{0, 1, 9, 16, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.lst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
