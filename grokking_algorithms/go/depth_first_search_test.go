package main

import "testing"

func Test_depthFirstSearch(t *testing.T) {
	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}

	type args struct {
		graph map[string][]string
		name  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found",
			args: args{
				graph: graph,
				name:  "you",
			},
			want: true,
		},
		{
			name: "not found",
			args: args{
				graph: graph,
				name:  "bob",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSellerDFS(tt.args.graph, tt.args.name, []string{}); got != tt.want {
				t.Errorf("depthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_depthFirstSearchV2(t *testing.T) {
	graph := map[string][]string{
		"A": {"E", "B", "C"},
		"E": {"I", "W"},
		"B": {"W", "D"},
		"C": {"U"},
		"I": {"Y", "U"},
		"W": {},
		"D": {},
		"Y": {"O"},
		"U": {"F"},
		"O": {},
		"F": {},
	}

	type args struct {
		graph map[string][]string
		v     string
		t     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found",
			args: args{
				graph: graph,
				v:     "A",
				t:     "F",
			},
			want: true,
		},
		{
			name: "not found",
			args: args{
				graph: graph,
				v:     "B",
				t:     "F",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := depthFirstSearch(tt.args.graph, tt.args.v, tt.args.t, []string{}); got != tt.want {
				t.Errorf("depthFirstSearchV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
