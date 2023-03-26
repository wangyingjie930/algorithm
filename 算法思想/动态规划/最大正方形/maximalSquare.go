/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc: https://leetcode.cn/problems/maximal-square/
**/

package 最大正方形

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	max := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				max = 1
			} else {
				dp[i][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] != '1' {
				continue
			}
			length := min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) //边长
			dp[i][j] = length + 1
			if dp[i][j]*dp[i][j] > max {
				max = dp[i][j] * dp[i][j] //面积
			}
		}
	}

	return max
}

func min(a ...int) int {
	m := a[0]
	for i := 1; i < len(a); i++ {
		if m > a[i] {
			m = a[i]
		}
	}
	return m
}
