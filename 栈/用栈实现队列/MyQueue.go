/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc: https://leetcode.cn/problems/implement-queue-using-stacks/
**/

package 用栈实现队列

// MyQueue
//  @Description: 思路就是使用两个栈的结构
type MyQueue struct {
	headStack []int
	tailStack []int
}

func Constructor() MyQueue {
	return MyQueue{headStack: []int{}, tailStack: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.tailStack = append(this.tailStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.headStack) == 0 {
		this.migrate()
	}
	ret := this.headStack[len(this.headStack)-1]
	this.headStack = this.headStack[0 : len(this.headStack)-1]
	return ret
}

func (this *MyQueue) migrate() {
	for len(this.tailStack) > 0 {
		val := this.tailStack[len(this.tailStack)-1]
		this.tailStack = this.tailStack[0 : len(this.tailStack)-1]
		this.headStack = append(this.headStack, val)
	}
}

func (this *MyQueue) Peek() int {
	if len(this.headStack) == 0 {
		this.migrate()
	}
	return this.headStack[len(this.headStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.tailStack) == 0 && len(this.headStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
