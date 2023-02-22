/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc: https://leetcode.cn/problems/number-of-islands/
**/

package 岛屿数量

func numIslands(grid [][]byte) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isIsland(grid, i, j) {
				num++
			}
		}
	}
	return num
}

func isIsland(grid [][]byte, i, j int) bool {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return false
	}
	if grid[i][j] == '0' {
		return false
	}

	grid[i][j] = '0'
	isIsland(grid, i-1, j)
	isIsland(grid, i+1, j)
	isIsland(grid, i, j-1)
	isIsland(grid, i, j+1)
	return true
}
