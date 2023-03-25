/**
  @author: wangyingjie
  @since: 2023/3/25
  @desc:
**/

package 最小栈最优解

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	minStack := Constructor()
	minStack.Push(1)
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	t.Log("top: ", minStack.Top())
	t.Log("min", minStack.GetMin())

	minStack.Pop()
	t.Log("min", minStack.GetMin())

	minStack.Pop()
	t.Log("min", minStack.GetMin())

	minStack.Pop()
	t.Log("min", minStack.GetMin())

	minStack.Pop()
	t.Log("min", minStack.GetMin())

	minStack.Pop()
	t.Log("min", minStack.GetMin())
}
