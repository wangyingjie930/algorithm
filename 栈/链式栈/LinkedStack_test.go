package 链式栈

import (
	"testing"
)

func TestLinkedStack_Pop(t *testing.T) {
	stack := NewLinkedStack()
	stack.Push("he")
	stack.Push("asd")
	stack.Push("sss")
	stack.Push("he")
	stack.Push("asd")
	stack.Push("sss")
	stack.Push("he")
	stack.Push("asd")
	stack.Push("sss")
	stack.Push("he")
	stack.Push("asd")
	stack.Push("sss")
	stack.Push("he")
	stack.Push("asd")
	stack.Push("sss")
	stack.Push("he")
	stack.Push("asd")
	stack.Push("sss")
	stack.Push("he")
	stack.Push("asd")
	stack.Push("sss")
	stack.Push("he")
	stack.Push("asd")
	stack.Push("sss")
	stack.Push("he")
	stack.Push("asd")
	stack.Push("sss")

	t.Logf("%+v", stack)
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
}
