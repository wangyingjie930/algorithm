/**
  @author: wangyingjie
  @since: 2023/3/28
  @desc: https://leetcode.cn/problems/basic-calculator-ii/
  https://labuladong.github.io/algo/di-san-zha-24031/jing-dian--a94a0/ru-he-shi--24fe4/
  掌握程度: 不熟
**/

package 基本计算器II

import "unicode"

// calculate
//  @Description: 1-12+3 转化为+1-12+3, 在栈中的表示为 [1, -12, 3], 全部之和相加即可
//  @param s
//  @return int
func calculate(s string) int {
	num := 0
	var stack []int
	sign := '+'

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			num = num*10 + int(s[i]-'0')
		}
		if (!unicode.IsDigit(rune(s[i])) && s[i] != ' ') || len(s) == i+1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * num
			case '/':
				stack[len(stack)-1] = stack[len(stack)-1] / num
			}
			sign = int32(s[i])
			num = 0
		}
	}

	sum := 0
	for i := 0; i < len(stack); i++ {
		sum += stack[i]
	}
	return sum
}
