/**
  @author: wangyingjie
  @since: 2023/2/26
  @desc: https://leetcode.cn/problems/longest-common-subsequence/description/
**/

package 最长公共子序列

import "math"

func longestCommonSubsequence(text1 string, text2 string) int {
	return dp(text1, text2)
}

func helper(text1 string, text2 string, i, j int) int {
	if i >= len(text1) || j >= len(text2) {
		return 0
	}
	if text1[i] == text2[j] {
		return helper(text1, text2, i+1, j+1) + 1
	}
	return int(math.Max(math.Max(
		float64(helper(text1, text2, i+1, j)),
		float64(helper(text1, text2, i, j+1)),
	), float64(helper(text1, text2, i+1, j+1))))
}

func dp(text1 string, text2 string) int {
	table := make([][]int, len(text2)+1)
	for i := 0; i <= len(text2); i++ {
		table[i] = make([]int, len(text1)+1)
	}
	for i := len(text2) - 1; i >= 0; i-- {
		for j := len(text1) - 1; j >= 0; j-- {
			if text1[j] == text2[i] {
				table[i][j] = table[i+1][j+1] + 1
			} else {
				table[i][j] = int(math.Max(math.Max(
					float64(table[i+1][j]),
					float64(table[i][j+1]),
				), float64(table[i+1][j+1])))
			}
		}
	}
	return table[0][0]
}
