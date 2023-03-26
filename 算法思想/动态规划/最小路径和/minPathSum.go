/**
  @author: wangyingjie
  @since: 2023/3/27
  @desc: https://leetcode.cn/problems/minimum-path-sum/description/
  简单, 不注释了, 一般有二维数组的动态规划都是以dp[i][j]的形式
**/

package 最小路径和

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
		if i == 0 {
			dp[0][0] = grid[0][0]
			for j := 1; j < len(grid[0]); j++ {
				dp[0][j] = dp[0][j-1] + grid[0][j]
			}
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
