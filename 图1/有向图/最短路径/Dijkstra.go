package 最短路径

import (
	"container/heap"
	"fmt"
	Graph "algorithm/图1/有向图"
	"math"
)

func Dijkstra(graph *Graph.Graph, start, end int) bool {
	queue := &NodePriorityQueue{Nodes: []*Graph.Node{}, Distance: map[*Graph.Node]int{}, hashedNodes: map[*Graph.Node]bool{}, path: map[*Graph.Node]*Graph.Node{}}
	heap.Init(queue)
	//将开始节点放入堆中
	heap.Push(queue, &item{Node: graph.Nodes[start], NewDistance: 0, Last: nil})
	isFind := false
	for queue.Len() > 0 {
		//此时弹出的是item对象, 包含的是弹出的节点和从开始节点到此节点的最短距离
		i := heap.Pop(queue).(*item)
		if i.Node.Val == end {
			isFind = true
			fmt.Println(i.NewDistance)
			break
		}
		for edge, _ := range i.Node.Edges {
			//将所有下个节点加入堆中
			//传入"从开始节点到当前节点的距离+当前节点到下节点的距离"
			//如果堆中存在下个节点则会以相比较小的距离为准
			heap.Push(queue, &item{Node: edge.To, NewDistance: edge.Weight + i.NewDistance, Last: edge.From})
		}
	}

	//打印路径
	for n1, n2 := range queue.path {
		fmt.Println(n1, n2)
	}
	return isFind
}

type NodePriorityQueue struct {
	// 待排序的节点
	Nodes []*Graph.Node
	// 每个节点从开始节点到此节点所经历的最短路径
	Distance map[*Graph.Node]int
	// 用于判断节点是否存在于堆中
	hashedNodes map[*Graph.Node]bool
	//path: 路径描述 值为上个节点
	path  map[*Graph.Node]*Graph.Node
}

// 由外部传入和弹出的数据结构
type item struct {
	// push需要push/pop的节点
	Node *Graph.Node
	// push时传入的需要判断的距离/pop弹出的最小距离
	NewDistance int
	//上个节点 (表示最短路径中的上个节点)
	Last *Graph.Node
}

func (q NodePriorityQueue) Len() int {
	return len(q.Nodes)
}

func (q NodePriorityQueue) Less(i, j int) bool {
	iNode := q.Nodes[i]
	jNode := q.Nodes[j]
	//以最短距离作为构建堆交换的条件
	return q.Distance[iNode] < q.Distance[jNode]
}

func (q *NodePriorityQueue) Swap(i, j int) {
	q.Nodes[i], q.Nodes[j] = q.Nodes[j], q.Nodes[i]
}

func (q *NodePriorityQueue) Push(x interface{}) {
	node := x.(*item)
	if _, find := q.hashedNodes[node.Node]; find {
		//如果堆中已经存在, 进行最小距离的判断, 并进行更新
		q.Distance[node.Node] = int(math.Min(float64(q.Distance[node.Node]), float64(node.NewDistance)))
		if q.Distance[node.Node] == node.NewDistance {
			//如果比原来堆的最短路径还小, 则记录下传入的到此节点的上个节点
			q.path[node.Node] = node.Last
		}
	}else {
		//如果不存在则正常插入
		q.Nodes = append(q.Nodes, node.Node)
		q.Distance[node.Node] = node.NewDistance
		q.hashedNodes[node.Node] = true
		q.path[node.Node] = node.Last
	}
}

func (q *NodePriorityQueue) Pop() interface{} {
	//弹出堆
	node := q.Nodes[len(q.Nodes) - 1]
	q.Nodes = q.Nodes[0:len(q.Nodes) - 1]
	distance := q.Distance[node]
	//对应的数据记录删除
	delete(q.Distance, node)
	delete(q.hashedNodes, node)
	return &item{Node: node, NewDistance: distance}
}
