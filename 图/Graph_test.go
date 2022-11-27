package Graph

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph([][]int{{1, 1, 2}, {2, 2, 3}, {3, 3, 4}})
	g.PrintGraph()
}

func TestGraph_Dfs(t *testing.T) {
	g := NewGraph([][]int{{1, 1, 2}, {2, 1, 3}, {3, 3, 4}})
	g.Dfs(g.Nodes[1])
}

func TestGraph_Bfs(t *testing.T) {
	g := NewGraph([][]int{{1, 1, 2}, {2, 1, 3}, {3, 3, 5}, {4, 3, 6}})
	g.Bfs(g.Nodes[1])
}

func TestGraph_SortedTopology(t *testing.T) {
	g := NewGraph([][]int{{1, 2, 3}, {2, 1, 3}, {3, 3, 5}, {4, 3, 6}})
	g.SortedTopology()
}

func TestGraph_KruskalMST(t *testing.T) {
	g := NewGraph([][]int{{33, 2, 3}, {11, 1, 3}, {22, 3, 5}, {3, 3, 6}, {1, 1, 2}, {2, 1, 6}})
	g.KruskalMST()
}

func TestGraph_PrimMst(t *testing.T) {
	g := NewGraph([][]int{{33, 2, 3}, {11, 1, 3}, {22, 3, 5}, {3, 3, 6}, {1, 1, 2}, {2, 1, 6}})
	g.PrimMst()
}
