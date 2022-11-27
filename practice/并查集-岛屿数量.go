package practice

func NumIslands(grid [][]byte) int {
	union := NewUnionFind(grid)

	/**
	{1,0,1,1,1},
	{1,0,1,0,1},
	{1,1,1,0,1},
	 */
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			grid[i][j] = 0
			//向当前岛屿的上下左右合并岛屿
			if j-1 >= 0 && grid[i][j-1] == 1 {
				union.Union(Element{Row: i, Col: j-1, Value: 1}, Element{Row: i, Col: j, Value: 1})
			}
			if i - 1 >= 0 && grid[i - 1][j] == 1 {
				union.Union(Element{Row: i - 1, Col: j, Value: 1}, Element{Row: i, Col: j, Value: 1})
			}
			if j + 1 < len(grid[0]) && grid[i][j + 1] == 1 {
				union.Union(Element{Row: i, Col: j+1, Value: 1}, Element{Row: i, Col: j, Value: 1})
			}
			if i + 1 < len(grid) && grid[i + 1][j] == 1 {
				union.Union(Element{Row: i + 1, Col: j, Value: 1}, Element{Row: i, Col: j, Value: 1})
			}
		}
	}

	return len(union.size)
}

type Element struct {
	Value byte
	Row int
	Col int
}

type UnionFind struct {
	parent map[Element]Element //子节点对应父节点
	size map[Element]int //父节点当前有多少个子节点
}

func NewUnionFind(grid [][]byte) UnionFind {
	var union = UnionFind{parent: map[Element]Element{}, size: map[Element]int{}}
	union.UnionFindSet(grid)
	return union
}

/**
初始化并查集集合
 */
func (u *UnionFind) UnionFindSet(grid [][]byte)  {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			e := Element{Value: grid[i][j], Row: i, Col: j}
			u.parent[e] = e
			u.size[e] = 1
		}
	}
}

/**
循环查找父节点, 直到找到头结点
 */
func (u *UnionFind) FindHead(e Element) Element {
	var stack []Element
	for u.parent[e] != e {
		e = u.parent[e]
		stack = append(stack, e)
	}
	//防止层级过高进行优化, 直接父级下面就只有一层子级
	for _, v := range stack {
		u.parent[v] = e
	}
	return e
}

func (u *UnionFind) IsSameSet(a Element, b Element) bool {
	return u.FindHead(a) == u.FindHead(b)
}

/**
合并集合, 加少的集合的父节点放入多的集合的父节点
合并集合的数量
删除被合并的变量
 */
func (u *UnionFind) Union(a Element, b Element) {
	aF := u.FindHead(a)
	bF := u.FindHead(b)
	if aF != bF {
		if u.size[aF] < u.size[bF] {
			u.parent[aF] = bF
			u.size[bF] = u.size[bF] + u.size[aF]
			delete(u.size, aF)
		}else {
			u.parent[bF] = aF
			u.size[aF] = u.size[bF] + u.size[aF]
			delete(u.size, bF)
		}
	}
}