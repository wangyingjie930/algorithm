package 岛屿数量

func numIslands(grid [][]byte) int {
	count := 0
	for i, iv := range grid {
		for j, jv := range iv {
			if jv != '1' {
				continue
			}
			//遇到1就进行走动, 走到整片区域都标记为2才接着下一轮
			bfs(&grid, i, j)
			//标记好多少轮, 就是有多少片相邻的岛
			count ++
		}
	}
	return count
}

func bfs(grid *[][]byte, i, j int) {
	//边界判断
	if i < 0 || i > len(*grid)-1 || j < 0 || j > len((*grid)[0])-1 {
		return
	}

	if (*grid)[i][j] != '1' {
		return
	}
	//进行标记, 以后不准走了
	(*grid)[i][j] = '2'

	//多个方向进行走动
	bfs(grid, i + 1, j)
	bfs(grid, i - 1, j)
	bfs(grid, i, j + 1)
	bfs(grid, i, j - 1)
}

