/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc: https://leetcode.cn/problems/decode-string/description/
**/

package 字符串解码

import "strings"

// decodeString
//  @Description: 解法1: 使用栈 https://leetcode-solution-leetcode-pp.gitbook.io/leetcode-solution/medium/394.decode-string
//  @param s
//  @return string
func decodeString(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			//弹出

			//得到需要重复的字符串
			repeatStr := ""
			for len(stack) > 0 {
				if stack[len(stack)-1] != '[' {
					repeatStr = string(stack[len(stack)-1]) + repeatStr
					stack = stack[:len(stack)-1]
				} else {
					stack = stack[:len(stack)-1]
					break
				}
			}

			//得到需要重复的次数
			repeatCount := 0
			digit := 1
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				repeatCount = int(stack[len(stack)-1]-'0')*digit + repeatCount
				stack = stack[:len(stack)-1]
				digit = 10 * digit
			}

			//展开对应的字符串, 并加入栈中
			stack = append(stack, strings.Repeat(repeatStr, repeatCount)...)
		} else {
			//插入
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

//todo: 使用递归的解法
