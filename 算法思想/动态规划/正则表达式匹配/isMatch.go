/**
  @author: wangyingjie
  @since: 2023/5/18
  @desc: https://leetcode.cn/problems/regular-expression-matching/
  @tag: unskilled
**/

package isMatch

//isMatch
//@Description:https://labuladong.unskilledgithub.io/algo/di-er-zhan-a01c6/yong-dong--63ceb/jing-dian--8d516/
//@param s string
//@param p string
//@return bool
func isMatch(s string, p string) bool {
	return helper(s, 0, p, 0)
}

func helper(s string, i int, p string, j int) bool {
	m, n := len(s), len(p)
	// base case
	if j == n {
		return i == m
	}
	if i == m {
		if (n-j)%2 == 1 {
			return false
		}
		for ; j+1 < n; j += 2 {
			if p[j+1] != '*' {
				return false
			}
		}
		return true
	}

	if s[i] == p[j] || p[j] == '.' {
		//当前对比的字符相同
		if j+1 < len(p) && p[j+1] == '*' {
			return helper(s, i, p, j+2) || //不进行匹配*
				helper(s, i+1, p, j) //尝试进行匹配*
		} else {
			return helper(s, i+1, p, j+1)
		}
	} else {
		//当前对比的字符不同
		if j+1 < len(p) && p[j+1] == '*' {
			return helper(s, i, p, j+2)
		} else {
			return false
		}
	}
}

//isMatchDp
//@Description:https://leetcode.cn/problems/regular-expression-matching/solutions/296114/shou-hui-tu-jie-wo-tai-nan-liao-by-hyj8/
//@param s string
//@param p string
//@return bool
func isMatchDp(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	//dp[i][j]: 表示s的前i个字符与p的前j个字符是否是匹配的
	//并且是从后往前进行匹配的
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			//当s为空串的时候, 且当前的p字符(下标为j-1)为'*', 则跳过'*'及'*'前面的字符从下标为 j-1-2进行匹配
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				//当前s和p的字符是相同的情况
				dp[i][j] = dp[i-1][j-1]
			} else {
				//不相同则看看比较的p是否为'*'
				if p[j-1] != '*' {
					//不是'*'直接判断不匹配
					dp[i][j] = false
				} else {
					//是'*', 还要看看'*'前面的字符是什么, 详细链接有画图
					if s[i-1] == p[j-2] || p[j-2] == '.' {
						dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
					} else {
						dp[i][j] = dp[i][j-2]
					}
				}
			}
		}
	}
	return dp[m][n]
}
