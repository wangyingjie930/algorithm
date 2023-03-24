/**
  @author: wangyingjie
  @since: 2023/3/25
  @desc: https://leetcode.cn/problems/min-stack/description/
**/

package 最小栈

import "math"

type MinStack struct {
	stack    []int //维护原有的数组
	minStack []int //辅助栈: 维护最小值数组
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, minStack: []int{math.MaxInt64}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	//拿到最小的数
	minVal := this.minStack[len(this.minStack)-1]
	if minVal > val {
		minVal = val
	}
	//放入辅助栈
	this.minStack = append(this.minStack, minVal)
}

func (this *MinStack) Pop() {
	if len(this.stack) <= 0 {
		return
	}

	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[0]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
