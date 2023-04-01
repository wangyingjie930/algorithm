/**
  @author: wangyingjie
  @since: 2023/4/1
  @desc: https://leetcode.cn/problems/egg-drop-with-2-eggs-and-n-floors/description/
  背吧
**/

package 鸡蛋掉落

import "math"

func twoEggDrop(n int) int {
	//dp[0][n]: 表示仅剩一个鸡蛋测试n层楼的最小操作次数
	//dp[1][n]: 表示剩两个鸡蛋测试n层楼的最小操作次数
	dp := make([][]int, 2)

	dp[0], dp[1] = make([]int, n+1), make([]int, n+1)
	dp[0][0], dp[1][0] = 0, 0

	for i := 1; i <= n; i++ {
		dp[0][i] = i //用1个鸡蛋的话只能层数从下往上一次进行测试
		dp[1][i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[1][i] = min(
				dp[1][i],
				max(
					dp[0][j-1]+1, //当在第j层碎了, 那么只能用1个去测试剩余的j-1层(dp[0][j-1])
					dp[1][i-j]+1, //当在第j层没碎, 那么用2个去测试剩余的i-j层(dp[1][i-j])
				),
			)
		}
	}

	return dp[1][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
