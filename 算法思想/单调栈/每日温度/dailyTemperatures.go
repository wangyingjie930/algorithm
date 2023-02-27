/**
  @author: wangyingjie
  @since: 2023/2/23
  @desc: https://leetcode.cn/problems/daily-temperatures/
**/

package 每日温度

func dailyTemperatures(temperatures []int) []int {
	var stack []int
	ret := make([]int, len(temperatures))

	//  栈中保持单调递减
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			ret[top] = i - top
		}
		stack = append(stack, i)
	}
	return ret
}
