package main

import (
	"reflect"
	"testing"
)

func Test_bellmanFord(t *testing.T) {
	g1 := NewBellmanFordGraph(5)
	g1.addEdge(0, 1, -1)
	g1.addEdge(0, 2, 4)
	g1.addEdge(1, 2, 3)
	g1.addEdge(1, 3, 2)
	g1.addEdge(1, 4, 2)
	g1.addEdge(3, 2, 5)
	g1.addEdge(3, 1, 1)
	g1.addEdge(4, 3, -3)

	g2 := NewBellmanFordGraph(4)
	g2.addEdge(0, 1, 1)
	g2.addEdge(1, 2, -1)
	g2.addEdge(2, 3, -1)
	g2.addEdge(3, 0, -1)

	type args struct {
		graph BellmanFordGraph
		start int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				graph: g1,
				start: 0,
			},
			want: []int{0, -1, 2, -2, 1},
		},
		{
			name: "2 negative cycle",
			args: args{
				graph: g2,
				start: 0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.graph.bellmanFord(tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bellmanFord() = %v, want %v", got, tt.want)
			}
		})
	}
}
