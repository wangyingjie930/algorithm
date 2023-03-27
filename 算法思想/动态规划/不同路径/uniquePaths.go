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

// uniqueOptimization
//  @Description: 优化成空间复杂度O(n), 面试重点考察
//  @param m
//  @param n
//  @return int
func uniqueOptimization(m int, n int) int {
	dp := make([]int, n)
	dp[0] = 1 //固定为1

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			//dp[j]表示的是上一层得到的数
			//dp[j-1]表示的是当前层左边得到的数
			//两者已加就可以得到走到当前位置的数
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}
