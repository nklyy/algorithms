package main

import (
	"reflect"
	"testing"
)

func Test_topologySort(t *testing.T) {
	graph1 := NewGraph(6)
	graph1.addEdge(5, 2)
	graph1.addEdge(5, 0)
	graph1.addEdge(4, 0)
	graph1.addEdge(4, 1)
	graph1.addEdge(2, 3)
	graph1.addEdge(3, 1)

	graph2 := NewGraph(7)
	graph2.addEdge(3, 2)
	graph2.addEdge(2, 1)
	graph2.addEdge(2, 0)
	graph2.addEdge(5, 4)
	graph2.addEdge(4, 0)
	graph2.addEdge(6, 0)

	graph3 := NewGraph(4)
	graph3.addEdge(0, 1)
	graph3.addEdge(0, 2)
	graph3.addEdge(1, 2)
	graph3.addEdge(2, 0)
	graph3.addEdge(2, 3)
	graph3.addEdge(3, 3)

	type args struct {
		graph TopologyGraph
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				graph: graph1,
			},
			want: []int{0, 1, 3, 2, 4, 5},
		},
		{
			name: "2",
			args: args{
				graph: graph2,
			},
			want: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "3",
			args: args{
				graph: graph3,
			},
			want: []int{3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.graph.topologySort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topologySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
