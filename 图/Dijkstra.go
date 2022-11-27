package Graph

import "math"

/**
从一个顶点到其余各顶点的最短路径算法
求出某个节点到到其余所有节点的最短路径

backend/图解算法/dijkstra算法.png
 */
func (g *Graph)Dijkstra() map[*Node]int {
	distanceMap := map[*Node]int{}
	var lockNodes []*Node
	for _, node := range g.Nodes {
		distanceMap[node] = math.MaxInt32
	}
	distanceMap[g.Nodes[1]] = 0

	for len(lockNodes) < len(g.Nodes) {
		minNode := getMinNode(distanceMap, lockNodes)
		for edge, _ := range minNode.Edges{
			distanceMap[edge.To] = int(math.Min(float64(distanceMap[edge.To]), float64(distanceMap[minNode] + edge.Weight)))
		}
		lockNodes = append(lockNodes, minNode)
	}
	return distanceMap
}

func getMinNode(distance map[*Node]int, lockNodes []*Node) *Node {
	return nil
}
