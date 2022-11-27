package 拓扑排序

import (
	"imooc-product/backend/图1/有向图"
	"testing"
)

func TestDfs(t *testing.T) {
	g := 有向图.NewGraph([][]int{{2, 3, 1}, {4, 5, 2}, {11, 4, 3}, {11, 2, 4}})
	Dfs(g)
}
