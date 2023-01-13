package 拓扑排序

import (
	"fmt"
	Graph "algorithm/图1/有向图"
)

func Kahn(graph *Graph.Graph)  {
	var queue []*Graph.Node
	inDegree := map[*Graph.Node]int{}
	//将所有入度为0的加入队列
	for _, node := range graph.Nodes {
		if node.In == 0 {
			queue = append(queue, node)
		}
		//统计每个节点对应的入度
		inDegree[node] = node.In
	}
	//从队列取数据
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		fmt.Println(n)

		//遍历下个节点数组将对应的入度都减1
		for node, _ := range n.Nexts {
			inDegree[node]--
			//入度为0的加入队列
			if inDegree[node] == 0 {
				queue = append(queue, node)
			}
		}
	}
}
