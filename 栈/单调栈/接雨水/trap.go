/**
  @author: wangyingjie
  @since: 2023/2/22
  @desc: https://leetcode.cn/problems/trapping-rain-water/
  @see: 单调栈讲解,九章算法: https://www.bilibili.com/video/BV1vu411r733/?vd_source=c0808a52efcf11b005f1df5936845dd2
**/

package 接雨水

import (
	"fmt"
	"math"
)

// trap
//  @Description: 单调栈思路
//  @param height
//  @return int
func trap(height []int) int {
	var stack []int
	res := 0

	//  单调栈: 以某跟柱子当做最小值, 向两边延伸
	//  由题意得, 要获得两边距离它最近的比它大的栈中需要保持栈中单调递减
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			//不满足单调递减, 弹出栈
			h := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if len(stack) == 0 {
				break
			}
			//栈弹出后的尾端为左边最近比它大的柱子, 右端则为刚刚尝试插入的i
			left, right := stack[len(stack)-1], i
			//h则为中间的柱子, 计算出空出来的长条面积 其中相差高度 = (左右两边取最短) - 中间高度
			res = res + (int(math.Min(float64(height[left]), float64(height[right])))-height[h])*(right-left-1)
			fmt.Printf("left: %d, right: %d, mid: %d, res: %d\n", left, right, h, res)
		}
		stack = append(stack, i)
	}
	return res
}
