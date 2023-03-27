/**
  @author: wangyingjie
  @since: 2023/3/28
  @desc: https://leetcode.cn/problems/basic-calculator-ii/
  https://labuladong.github.io/algo/di-san-zha-24031/jing-dian--a94a0/ru-he-shi--24fe4/
**/

package 基本计算器II

import "unicode"

// calculate
//  @Description: 1-12+3 转化为+1-12+3, 在栈中的表示为 [1, -12, 3], 全部之和相加即可
//  @param s
//  @return int
func calculate(s string) int {
	var stack []int
	num := 0
	var sign byte = '+' //起始默认为正数开头

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			num = num*10 + int(s[i]-'0')
		}

		//i == len(s)-1: 循环到最后的时候, 没有时机插入栈了, 所以在这次循环中一起插入栈里
		if (!unicode.IsDigit(rune(s[i])) && s[i] != ' ') || i == len(s)-1 {
			//新运算符来临, 是处理上一次的num和sign的
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				num = stack[len(stack)-1] * num
				stack[len(stack)-1] = num
			case '/':
				num = stack[len(stack)-1] / num
				stack[len(stack)-1] = num
			}
			sign = s[i] //将新符号赋值, 供下一个新符号来时可以判断是否为正负数
			num = 0     //数字清零
		}
	}

	ret := 0
	for i := 0; i < len(stack); i++ {
		ret = ret + stack[i]
	}
	return ret
}
