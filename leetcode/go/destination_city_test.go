package main

import (
	"testing"
)

func TestDestCity(t *testing.T) {
	tests := []struct {
		paths [][]string
		want  string
	}{
		{[][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}, "Sao Paulo"},
		{[][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}, "A"},
	}

	for _, tt := range tests {
		got := destCity(tt.paths)
		if got != tt.want {
			t.Errorf("destCity(%v) = %v; want %v", tt.paths, got, tt.want)
		}
	}
}
