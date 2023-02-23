/**
  @author: wangyingjie
  @since: 2023/2/23
  @desc: https://leetcode.cn/problems/climbing-stairs/
**/

package 爬楼梯

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 0 {
		return 0
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsDp(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
