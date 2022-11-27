package Graph

import (
	"fmt"
	"sort"
)

type sortEdges []*Edge

func (edges sortEdges) Len() int {
	return len(edges)
}
func (edges sortEdges) Less(i, j int) bool {
	return edges[i].Weight < edges[j].Weight
}
func (edges sortEdges) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}


type unionFind struct {
	parent map[*Node]*Node //子节点对应父节点
	size map[*Node]int //父节点当前有多少个子节点
}
func newUnionFind(nodes map[int]*Node) *unionFind  {
	unionFind := &unionFind{parent: map[*Node]*Node{}, size: map[*Node]int{}}
	for _, node := range nodes {
		unionFind.parent[node] = node
		unionFind.size[node] = 1
	}
	return unionFind
}
func (u *unionFind)isSameSer(a, b *Node) bool  {
	return u.findHead(a) == u.findHead(b)
}
func (u *unionFind)findHead(node *Node) *Node  {
	for node != u.parent[node] {
		node = u.parent[node]
	}
	return node
}
func (u *unionFind)union(a, b *Node)  {
	headA := u.findHead(a)
	headB := u.findHead(b)
	if headA == headB {
		return
	}
	if u.size[headA] < u.size[headB] {
		u.parent[headA] = headB
		u.size[headB] = u.size[headA] + u.size[headB]
		delete(u.size, headA)
	}else {
		u.parent[headB] = headA
		u.size[headA] = u.size[headB] + u.size[headA]
		delete(u.size, headB)
	}
}


/**
最小生成树, K算法 - 无向图
1. 对每一条边进行从小到大的排序
2. 选择最小的边, 观察是否会形成环, 如果会那就换下一条边, 不会则选择

那么如何判断是否形成环呢? 使用的是并查集
*/
func (g *Graph) KruskalMST() {
	//1. 按照边的权重排序
	var sortEdges sortEdges
	for edge, _ := range g.Edges {
		sortEdges = append(sortEdges, edge)
	}
	sort.Sort(sortEdges)
	union := newUnionFind(g.Nodes)
	//初始化并查集
	for i := 0; i < len(sortEdges); i ++ {
		edge := sortEdges[i]
		if sameSet := union.isSameSer(edge.From, edge.To); sameSet {
			continue
		}
		union.union(edge.From, edge.To)
		fmt.Println(edge)
	}
}
