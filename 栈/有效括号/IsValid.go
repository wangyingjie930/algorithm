// Package 有效括号
// @desc https://leetcode.cn/problems/valid-parentheses/
package 有效括号

func IsValid(str string) bool {
	var stack []byte
	var rightToLeft = map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}

	for i := 0; i < len(str); i++ {
		if v, exist := rightToLeft[str[i]]; exist {
			if len(stack) == 0 {
				return false
			}
			popNum := stack[len(stack)-1]
			if popNum != v {
				return false
			} else {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, str[i])
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
