package practice

// ../图解算法/最长有效括号.png

func LongestValidParentheses(s string) int {
	dp := make([]int, len(s))
	max := 0
	for k, v := range s {
		if v == '(' || k - 1 < 0 {
			continue
		}
		pre := k - dp[k - 1] - 1
		if pre >= 0 && s[pre] == '(' {
			predata := 0
			if pre - 1 >= 0  {
				predata = dp[pre - 1]
			}
			dp[k] = predata + dp[k - 1] + 2
		}
		if max < dp[k] {
			max = dp[k]
		}
	}
	return max
}
