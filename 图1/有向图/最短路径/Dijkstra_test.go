package 最短路径

import (
	"algorithm/图1/有向图"
	"testing"
)

func TestDijkstra(t *testing.T) {
	//11->2-->3
	//11->4->5
	g := 有向图.NewGraph([][]int{{0, 1, 10}, {1, 2, 15}, {2, 5, 5},
		{1, 3, 2}, {3, 2, 1}, {3, 5, 12}, {0, 4, 15}, {4, 5, 10}})
	Dijkstra(g, 0, 5)
}
