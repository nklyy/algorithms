package main

import (
	"reflect"
	"testing"
)

func Test_dijkstra(t *testing.T) {
	g1 := NewDijkstraGraph(4)
	g1.addEdge(0, 1, 5)
	g1.addEdge(1, 2, 6)
	g1.addEdge(2, 3, 2)
	g1.addEdge(0, 2, 15)

	g2 := NewDijkstraGraph(4)
	g2.addEdge(0, 1, 24)
	g2.addEdge(0, 3, 20)
	g2.addEdge(2, 0, 3)
	g2.addEdge(3, 2, 12)

	g3 := NewDijkstraGraph(4)
	g3.addEdge(0, 1, 6)
	g3.addEdge(0, 2, 2)
	g3.addEdge(1, 3, 1)
	g3.addEdge(2, 1, 3)
	g3.addEdge(2, 3, 5)

	g4 := NewDijkstraGraph(9)
	g4.addEdge(0, 1, 4)
	g4.addEdge(0, 7, 8)
	g4.addEdge(1, 2, 8)
	g4.addEdge(1, 7, 11)
	g4.addEdge(2, 3, 7)
	g4.addEdge(2, 8, 2)
	g4.addEdge(2, 5, 4)
	g4.addEdge(3, 4, 9)
	g4.addEdge(3, 5, 14)
	g4.addEdge(4, 5, 10)
	g4.addEdge(5, 6, 2)
	g4.addEdge(6, 7, 1)
	g4.addEdge(6, 8, 6)
	g4.addEdge(7, 8, 7)

	type args struct {
		graph DijkstraGraph
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
			want: []int{0, 5, 11, 13},
		},
		{
			name: "2",
			args: args{
				graph: g2,
				start: 0,
			},
			want: []int{0, 24, 3, 15},
		},
		{
			name: "3",
			args: args{
				graph: g3,
				start: 0,
			},
			want: []int{0, 5, 2, 6},
		},
		{
			name: "4",
			args: args{
				graph: g4,
				start: 0,
			},
			want: []int{0, 4, 12, 19, 21, 11, 9, 8, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.graph.dijkstra(tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dijkstra() = %v, want %v", got, tt.want)
			}
		})
	}
}
