package 拓扑排序

import (
	"algorithm/图1/有向图"
	"testing"
)

func TestKahn(t *testing.T) {
	g := 有向图.NewGraph([][]int{{2, 3, 1}, {4, 5, 2}, {11, 4, 3}, {11, 2, 4}})
	Kahn(g)
}
