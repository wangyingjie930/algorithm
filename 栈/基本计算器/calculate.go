/**
  @author: wangyingjie
  @since: 2023/5/18
  @desc: https://leetcode.cn/problems/basic-calculator/description/
**/

package calculate

func calculate(s1 string) int {
	s := []byte(s1)
	return helper(&s) //涉及到修改切片的长度, 使用指针比较好
}

func helper(s *[]byte) int {
	num := 0
	var sign byte = '+'
	var stack []int

	for len(*s) > 0 {
		cur := (*s)[0]
		*s = (*s)[1:]
		if cur >= '0' && cur <= '9' {
			num = num*10 + int(cur-'0')
		}
		if cur == '(' { //遇到'('则进入递归
			num = helper(s)
		}
		if (!(cur >= '0' && cur <= '9') && cur != ' ') || len(*s) == 0 {
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
			num = 0
			sign = cur
		}
		if cur == ')' { //放在这个位置是为了让')'结尾的数可以进入stack
			break
		}
	}

	sum := 0
	for i := 0; i < len(stack); i++ {
		sum += stack[i]
	}
	return sum
}
