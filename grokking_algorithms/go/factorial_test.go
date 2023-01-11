package main

import "testing"

func Test_factorial(t *testing.T) {
	type args struct {
		x uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "test 1",
			args: args{
				x: 3,
			},
			want: 6,
		},
		{
			name: "test 2",
			args: args{
				x: 5,
			},
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.args.x); got != tt.want {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorialLoop(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				x: 3,
			},
			want: 6,
		},
		{
			name: "test 2",
			args: args{
				x: 5,
			},
			want: 120,
		},
		{
			name: "test 3",
			args: args{
				x: 10,
			},
			want: 3628800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialLoop(tt.args.x); got != tt.want {
				t.Errorf("factorialLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
