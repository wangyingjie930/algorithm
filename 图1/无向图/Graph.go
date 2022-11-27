package 无向图

import "fmt"

type void struct {}

// Graph 无向图
type Graph struct {
	Nodes map[int]*Node
	Edges map[*Edge]void
}

func NewGraph(matrix [][]int) *Graph {
	grape := &Graph{Nodes: map[int]*Node{}, Edges: map[*Edge]void{}}
	for _, values := range matrix {
		from := values[0]
		to := values[1]
		weight := values[2]

		if _, find := grape.Nodes[from]; !find {
			grape.Nodes[from] = NewNode(from)
		}
		if _, find := grape.Nodes[to]; !find {
			grape.Nodes[to] = NewNode(to)
		}

		fromEdge := NewEdge(grape.Nodes[from], grape.Nodes[to], weight)
		toEdge := NewEdge(grape.Nodes[to], grape.Nodes[from], weight)

		grape.Nodes[from].In++
		grape.Nodes[to].In++
		grape.Nodes[from].Out++
		grape.Nodes[to].Out++

		grape.Nodes[from].Nexts[grape.Nodes[to]] = void{}
		grape.Nodes[to].Nexts[grape.Nodes[from]] = void{}
		grape.Nodes[from].Edges[fromEdge] = void{}
		grape.Nodes[to].Edges[toEdge] = void{}

		grape.Edges[fromEdge] = void{}
		grape.Edges[toEdge] = void{}
	}
	return grape
}

type Node struct {
	Val int
	In int
	Out int
	Nexts map[*Node]void
	Edges map[*Edge]void
}

func NewNode(val int) *Node {
	return &Node{Val: val, In: 0, Out: 0, Nexts: map[*Node]void{}, Edges: map[*Edge]void{}}
}

type Edge struct {
	From *Node
	To *Node
	Weight int
}

func NewEdge(from *Node, to *Node, weight int) *Edge {
	return &Edge{From: from, To: to, Weight: weight}
}

func (node *Node) String() string {
	return fmt.Sprintf("编号: %+v, In: %+v, Out: %+v", node.Val, node.In, node.Out)
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
