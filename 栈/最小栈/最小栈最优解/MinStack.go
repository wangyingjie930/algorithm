/**
  @author: wangyingjie
  @since: 2023/3/25
  @desc: https://leetcode.cn/problems/min-stack/description/
**/

package 最小栈最优解

type MinStack struct {
	stack []int //栈中维护的是val-min(入栈前的min)的差值
	min   int   //存放的是最小值
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, min: 0}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.min = val
	} else {
		diff := val - this.min
		this.stack = append(this.stack, diff)
		if diff < 0 {
			this.min = val
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) <= 0 {
		return
	}

	diff := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if diff < 0 {
		//说明此时pop出的是最小值, 并且最小值已经放在min中, 即: val = 当前的min
		//需要做的是还原上一个最小值, 因为之前是diff = val - 上一个min, 那么上一个min = val - diff
		this.min = this.min - diff
	}
	if len(this.stack) == 0 {
		this.min = 0
	}
}

func (this *MinStack) Top() int {
	diff := this.stack[len(this.stack)-1]
	if diff < 0 {
		return this.min
	} else {
		return diff + this.min
	}
}

func (this *MinStack) GetMin() int {
	return this.min
}
