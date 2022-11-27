package Graph

import (
	heap2 "container/heap"
	"fmt"
)

/**
1. 随机获取一个节点
2. 将该节点的所有边放入优先队列, 找出这个节点对应的最小的边
3. 跳到这个节点, 将它保存到集合里面, 将该节点的所有边存入优先队列
4. 弹出最小的边, 如果该节点不在集合则可以按照以上步骤, 如果该节点存在集合中, 继续弹出边
 */
type heapEdges []*Edge

func (h heapEdges) Len() int {
	return len(h)
}

func (h heapEdges) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h heapEdges) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapEdges) Push(x interface{}) {
	*h = append(*h, x.(*Edge))
}

func (h *heapEdges) Pop() interface{} {
	last := (*h)[h.Len() - 1]
	*h = (*h)[0:h.Len() - 1]
	return last
}

func (g *Graph) PrimMst() {
	heapEdge := &heapEdges{}
	heap2.Init(heapEdge)
	node := g.Nodes[6]
	for edge, _ := range node.Edges {
		heap2.Push(heapEdge, edge)
	}

	nodeSet := map[*Node]bool{}
	nodeSet[node] = true

	for heapEdge.Len() > 0 {
		primEdge := heap2.Pop(heapEdge)
		nextNode := primEdge.(*Edge).To
		if nextNode == node {
			nextNode = primEdge.(*Edge).From
		}
		if _, find := nodeSet[nextNode]; !find {
			node = nextNode
			nodeSet[nextNode] = true
			fmt.Println(primEdge)
			for edge, _ := range nextNode.Edges {
				heap2.Push(heapEdge, edge)
			}
		}
	}
}
