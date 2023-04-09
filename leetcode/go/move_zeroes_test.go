package main

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
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
				lst: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "2",
			args: args{
				lst: []int{0},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.lst)

			if !reflect.DeepEqual(tt.args.lst, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", tt.args.lst, tt.want)
			}
		})
	}
}
