package Graph

import (
	"fmt"
)

// Node 节点
type Node struct {
	Value int
	In int // 入度
	Out int // 出度
	Nexts map[*Node]bool //与该节点相邻的所有点集
	Edges map[*Edge]bool //与该节点相连的所有边集
}

func newNode(value int) *Node {
	return &Node{Value: value, In: 0, Out: 0, Nexts: map[*Node]bool{}, Edges: map[*Edge]bool{}}
}

// Edge 边
type Edge struct {
	Weight int
	From *Node
	To *Node
}

func newEdge(weight int, from, to *Node) *Edge {
	return &Edge{Weight: weight, From: from, To: to}
}

// Graph 图
type Graph struct {
	Nodes map[int]*Node //点集
	Edges map[*Edge]bool //边集
}

func NewGraph(matrix [][]int) *Graph {
	graph := &Graph{Nodes: map[int]*Node{}, Edges: map[*Edge]bool{}}
	for _, v := range matrix {
		weight := v[0]
		from := v[1]
		to := v[2]

		if _, find := graph.Nodes[from]; !find {
			graph.Nodes[from] = newNode(from)
		}
		if _, find := graph.Nodes[to]; !find {
			graph.Nodes[to] = newNode(to)
		}

		fromNode := graph.Nodes[from]
		toNode := graph.Nodes[to]
		edge := newEdge(weight, fromNode, toNode)
		graph.Edges[edge] = true

		fromNode.Out ++
		fromNode.Nexts[toNode] = true
		fromNode.Edges[edge] = true

		toNode.In ++
		toNode.Edges[edge] = true
	}
	return graph
}

func (node *Node) String() string {
	return fmt.Sprintf("编号: %+v, In: %+v, Out: %+v", node.Value, node.In, node.Out)
}

func (edge *Edge) String() string {
	return fmt.Sprintf("Weight: %d, From: (%s), To: (%s)", edge.Weight, edge.From, edge.To)
}

func (g *Graph) PrintGraph() {
	for _, node := range g.Nodes {
		fmt.Println(node)
		for node, _ := range node.Nexts{
			fmt.Printf("next_nodes: %s\n", node)
		}
		for edge, _ := range node.Edges{
			fmt.Printf("next_edges: %s\n", edge)
		}
		fmt.Println(" ")
	}
}

/**
图的广度优先遍历
1, 利用队列实现
2, 从源节点开始依次按照宽度进队列，然后弹出
3, 每弹出一个点，把该节点所有没有进过队列的邻接点放入队列
4, 直到队列变空
*/
func (g *Graph) Bfs(node *Node)  {
	fmt.Printf("%+v\n", *node)
	set := map[*Node]bool{node: true}
	queue := []*Node{node}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for v := range cur.Nexts {
			if _, find := set[v]; find {
				continue
			}
			fmt.Printf("%+v\n", *v)
			queue = append(queue, v)
		}
	}
}

/**
深度优先遍历
1，利用栈实现
2，从源节点开始把节点按照深度放入栈，然后弹出
3，每弹出一个点，把该节点下一个没有进过栈的邻接点放入栈
4. 直到栈变空
 */
func (g *Graph) Dfs(node *Node)  {
	fmt.Printf("%+v\n", *node)
	set := map[*Node]bool{node: true}
	stack := []*Node{node}
	for len(stack) > 0 {
		tmp := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		for v := range tmp.Nexts {
			if _, find := set[v]; find {
				continue
			}
			stack = append(stack, tmp, v)
			set[v] = true
			fmt.Printf("%+v\n", *v)
			break
		}
	}
}

/**
拓扑排序
1. 维护一个入度为0节点数组
2. 擦除入度为0及它产生的连接, 将入度为0的节点插入数组
3. 循环获得入度为0的节点, 擦除影响....
 */
func (g *Graph) SortedTopology() {
	var zeroIn []*Node
	for _, node := range g.Nodes {
		if node.In == 0 {
			zeroIn = append(zeroIn, node)
		}
	}
	for len(zeroIn) > 0 {
		zeroNode := zeroIn[0]
		zeroIn = zeroIn[1:]
		fmt.Printf("%+v\n", zeroNode)
		for node, _ := range zeroNode.Nexts {
			node.In --
			if node.In == 0 {
				zeroIn = append(zeroIn, node)
			}
		}
	}
}
