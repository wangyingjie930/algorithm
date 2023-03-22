/**
  @author: wangyingjie
  @since: 2023/3/22
  @desc: https://leetcode.cn/problems/generate-parentheses/
**/

package 括号生成

func generateParenthesis(n int) []string {
	cur := "(" //符合条件的必定是以(开头
	var ret []string
	helper(cur, n-1, n, &ret)
	return ret
}

func helper(cur string, leftNum int, rightNum int, ret *[]string) {
	if leftNum == 0 && rightNum == 0 {
		if isValid(cur) {
			*ret = append(*ret, cur)
		}
		return
	} else if leftNum < 0 || rightNum < 0 {
		return
	}

	cur = cur + "("
	helper(cur, leftNum-1, rightNum, ret)
	cur = cur[0 : len(cur)-1]

	cur = cur + ")"
	helper(cur, leftNum, rightNum-1, ret)
	cur = cur[0 : len(cur)-1]
}

var stack []byte

func isValid(s string) bool {
	stack = []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else if s[i] == ')' {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[0 : len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		}
	}
	return len(stack) == 0
}
