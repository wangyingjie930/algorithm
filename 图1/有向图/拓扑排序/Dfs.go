package 拓扑排序

import Graph "algorithm/图1/有向图"

// 构建逆序的图 错的
func inverseGraph(graph *Graph.Graph) *Graph.Graph {
	inverseGraph := &Graph.Graph{Nodes: map[int]*Graph.Node{}, Edges: map[*Graph.Edge]Graph.Void{}}
	for _, node := range graph.Nodes {
		if _, find := inverseGraph.Nodes[node.Val]; !find {
			inverseGraph.Nodes[node.Val] = Graph.NewNode(node.Val)
			inverseGraph.Nodes[node.Val].In = node.Out
			inverseGraph.Nodes[node.Val].Out = node.In
		}
		newNode := inverseGraph.Nodes[node.Val]

		for edge, _ := range node.Edges {
			newEdge := Graph.NewEdge(edge.To, edge.From, edge.Weight)
			newNode.Edges[newEdge] = struct{}{}

			if _, find := inverseGraph.Nodes[edge.From.Val]; !find {
				inverseGraph.Nodes[edge.From.Val] = Graph.NewNode(edge.From.Val)
				inverseGraph.Nodes[edge.From.Val].In = edge.From.Out
				inverseGraph.Nodes[edge.From.Val].Out = edge.To.In
			}

			newNode.Nexts[inverseGraph.Nodes[edge.From.Val]] = struct{}{}
			inverseGraph.Edges[newEdge] = struct{}{}
		}
	}
	return inverseGraph
}

func Dfs(graph *Graph.Graph)  {
	inverse := inverseGraph(graph)
	inverse.PrintGraph()
}
