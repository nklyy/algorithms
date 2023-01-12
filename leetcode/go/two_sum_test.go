package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		lst    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				lst:    []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "test 2",
			args: args{
				lst:    []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "test 3",
			args: args{
				lst:    []int{2, 11, 7, 15},
				target: 9,
			},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumV1(tt.args.lst, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumV1() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumV2(tt.args.lst, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumV2() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumV3(tt.args.lst, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
