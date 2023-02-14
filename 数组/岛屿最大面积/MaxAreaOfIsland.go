/**
  @author: wangyingjie
  @since: 2023/2/14
  @desc: https://leetcode.cn/problems/max-area-of-island/description/
**/

package 岛屿最大面积

import "math"

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			//遍历获取每个元素的面积
			max = int(math.Max(float64(getArea(grid, i, j)), float64(max)))
		}
	}
	return max
}

// getArea
//  @Description: 获取每个元素向外延伸之后的面积
//  @param grid
//  @param i
//  @param j
//  @return int
func getArea(grid [][]int, i, j int) int {
	//边界
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	//每访问一个修改标记一个, 下次就不用再进入判断
	grid[i][j] = 0

	//获取向外延伸的面积
	return getArea(grid, i-1, j) + getArea(grid, i+1, j) + getArea(grid, i, j-1) + getArea(grid, i, j+1) + 1
}
