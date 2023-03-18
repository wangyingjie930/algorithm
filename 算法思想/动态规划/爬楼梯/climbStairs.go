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

//不能爬到7及7的倍数
func climbStairsDpSeven(n int) int {
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
		if i%7 == 0 {
			dp[i] = 0
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}

	return dp[n]
}

//要求输出每条路径
func climbStairsBacktrack(n int) [][]int {
	var cur []int
	var ret [][]int
	backtrackHelper(cur, n, &ret)
	return ret
}

func backtrackHelper(cur []int, n int, ret *[][]int) {
	if n < 0 {
		return
	}
	if n == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
	}

	for i := 1; i <= 2; i++ {
		cur = append(cur, i)
		backtrackHelper(cur, n-i, ret)
		cur = cur[0 : len(cur)-1]
	}
}
