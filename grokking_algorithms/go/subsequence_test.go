package main

import (
	"reflect"
	"testing"
)

func Test_subsequence(t *testing.T) {
	type args struct {
		word_a string
		word_b string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				word_a: "fish",
				word_b: "fosh",
			},
			want: [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 2, 2}, {1, 1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsequence(tt.args.word_a, tt.args.word_b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
