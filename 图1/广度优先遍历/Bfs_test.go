package 广度优先遍历

import (
	"imooc-product/backend/图1/无向图"
	"testing"
)

func TestBfs(t *testing.T) {
	g := 无向图.NewGraph([][]int{{2, 3, 1}, {4, 5, 2}, {11, 4, 3}, {11, 2, 4}})
	Bfs(g, 11)
}
