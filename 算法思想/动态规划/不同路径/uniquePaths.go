/**
  @author: wangyingjie
  @since: 2023/3/28
  @desc: https://leetcode.cn/problems/unique-paths/
**/

package 不同路径

// uniquePaths
//  @Description: 标准动态规划不多逼逼
//  @param m
//  @param n
//  @return int
func uniquePaths(m int, n int) int {
	//dp[i][j] = dp[i-1][j] + dp[i][j-1]
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				//初始化
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
