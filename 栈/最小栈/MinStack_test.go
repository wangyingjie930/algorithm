/**
  @author: wangyingjie
  @since: 2023/3/25
  @desc:
**/

package 最小栈

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	minStack := Constructor()
	minStack.Push(1)
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	t.Log(minStack.Top())
	t.Log(minStack.GetMin())

	minStack.Pop()
	t.Log(minStack.GetMin())

	minStack.Pop()
	t.Log(minStack.GetMin())

	minStack.Pop()
	t.Log(minStack.GetMin())

	minStack.Pop()
	t.Log(minStack.GetMin())

	minStack.Pop()
	t.Log(minStack.GetMin())
}
