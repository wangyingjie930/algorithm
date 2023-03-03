/**
  @author: wangyingjie
  @since: 2023/3/3
  @desc: https://leetcode.cn/problems/minimum-add-to-make-parentheses-valid/
**/

package 使括号有效的最少添加

// minAddToMakeValid
//  @Description: 用的是栈, 左右括号相互抵消, 剩下的待匹配的就是最少需要添加的
//  @param s
//  @return int
func minAddToMakeValid(s string) int {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && s[i] == ')' {
			popNum := stack[len(stack)-1]
			if popNum == '(' {
				//除了左右括号匹配, 其他情况都压入栈中
				stack = stack[0 : len(stack)-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}
	//统计栈中还剩多少待匹配的字符
	return len(stack)
}
