package main

import "testing"

func Test_depthFirstSearch(t *testing.T) {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

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
			if got := depthFirstSearch(tt.args.graph, tt.args.name); got != tt.want {
				t.Errorf("depthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_depthFirstSearchV2(t *testing.T) {
	graph := make(map[string][]string)
	graph["A"] = []string{"E", "B", "C"}
	graph["E"] = []string{"I", "W"}
	graph["B"] = []string{"W", "D"}
	graph["C"] = []string{"U"}
	graph["I"] = []string{"Y", "U"}
	graph["W"] = []string{}
	graph["D"] = []string{}
	graph["Y"] = []string{"O"}
	graph["U"] = []string{"F"}
	graph["O"] = []string{}
	graph["F"] = []string{}

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
			if got := depthFirstSearchV2(tt.args.graph, tt.args.v, tt.args.t); got != tt.want {
				t.Errorf("depthFirstSearchV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
