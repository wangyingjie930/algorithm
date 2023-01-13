package 深度优先遍历

import (
	"fmt"
	"algorithm/图1/无向图"
)

// Dfs 递归写法
func Dfs(graph *无向图.Graph, start int, visited map[int]bool)  {
	if _, find := visited[start]; find {
		return
	}

	visited[start] = true
	fmt.Println(graph.Nodes[start])
	for node, _ := range graph.Nodes[start].Nexts {
		Dfs(graph, node.Val, visited)
	}
}

// Dfs1 使用栈
func Dfs1(graph *无向图.Graph, start int)  {
	// 保存访问过的节点
	visited := map[*无向图.Node]bool{graph.Nodes[start]: true}
	//深度优先用栈
	stack := []*无向图.Node{graph.Nodes[start]}

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[0:len(stack) - 1]
		fmt.Println(node)
		visited[node] = true

		//将所有子节点放入堆栈中
		for n, _ := range node.Nexts {
			if _, find := visited[n]; !find {
				stack = append(stack, n)
			}
		}
	}
}