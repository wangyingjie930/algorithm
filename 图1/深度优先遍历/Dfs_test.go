package 深度优先遍历

import (
	"fmt"
	"algorithm/图1/无向图"
	"testing"
)

func TestDfs(t *testing.T) {
	g := 无向图.NewGraph([][]int{{1, 2, 1}, {2, 3, 2}, {3, 4, 3}})
	Dfs(g, 2, map[int]bool{})

	fmt.Println("-------------")
	g = 无向图.NewGraph([][]int{{2, 3, 1}, {4, 5, 2}, {11, 4, 3}, {11, 2, 4}})
	Dfs(g, 2, map[int]bool{})

	fmt.Println("-------------")
	Dfs1(g, 2)
}
