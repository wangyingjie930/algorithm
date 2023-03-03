/**
  @author: wangyingjie
  @since: 2023/3/3
  @desc: https://leetcode.cn/problems/longest-valid-parentheses/description/
  @label: 需回看
**/

package 最长有效括号

// longestValidParentheses
//  @Description:
//  @param s
//  @return int
//	todo: 待掌握3种解法
func longestValidParentheses(s string) int {
	//定义dp的含义: dp[i]表示以i作为结尾的最长有效括号
	//状态, 选择....
	if len(s) == 0 {
		return 0
	}

	var stack []int
	dp := make([]int, len(s))
	maxLen := 0
	dp[0] = 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			//压入的是字符串的索引
			stack = append(stack, i)
			dp[i] = 0
		} else {
			if len(stack) > 0 {
				//拿到对应的字符串索引
				popIndex := stack[len(stack)-1]
				//依赖于弹出匹配字符的前一个字符的最长有效括号长度
				if s[popIndex] == '(' {
					stack = stack[0 : len(stack)-1]
					if popIndex-1 < 0 {
						dp[i] = i - popIndex + 1
					} else {
						//前一个字符的最长有效括号长度+弹出字符距当前字符的距离
						dp[i] = dp[popIndex-1] + (i - popIndex + 1)
					}
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}
