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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := depthFirstSearch(tt.args.graph, tt.args.name); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
