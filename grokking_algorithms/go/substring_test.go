package main

import (
	"reflect"
	"testing"
)

func Test_substring(t *testing.T) {
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
				word_b: "hish",
			},
			want: [][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 2, 0}, {0, 0, 0, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substring(tt.args.word_a, tt.args.word_b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("substring() = %v, want %v", got, tt.want)
			}
		})
	}
}
