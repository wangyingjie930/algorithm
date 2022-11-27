/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合

输入：s = "()[]{}"
输出：true
 */
package 有效括号

var rightToLeft = map[byte]byte{
	']': '[',
	'}': '{',
	')': '(',
}

func IsValid(str string) bool {
	var stack []byte
	for k, _ := range str {
		if v1, found := rightToLeft[str[k]]; found {
			index := len(stack) - 1
			if index < 0 {
				return false
			}
			j := stack[index]
			if j == v1 {
				stack = stack[:index]
			}else {
				stack = append(stack, str[k])
			}
		}else {
			stack = append(stack, str[k])
		}
	}
	return len(stack) == 0
}