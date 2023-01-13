package 广度优先遍历

import (
	"fmt"
	"algorithm/图1/无向图"
)

func Bfs(graph *无向图.Graph, start int)  {
	//广度优先用队列
	queue := []*无向图.Node{graph.Nodes[start]}
	//访问标记
	visited := map[*无向图.Node]bool{graph.Nodes[start]: true}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node)
		visited[node] = true

		//遍历所有子节点, 加入队列
		for n, _ := range node.Nexts {
			if _, find := visited[n]; !find {
				queue = append(queue, n)
			}
		}
	}
}
