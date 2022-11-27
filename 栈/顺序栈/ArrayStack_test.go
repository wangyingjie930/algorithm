package 顺序栈

import "testing"

func TestArrayStack_push(t *testing.T) {
	stack := NewArrayStack(5)
	stack.push("he")
	stack.push("asd")
	stack.push("sss")
	stack.push("he")
	stack.push("asd")
	stack.push("sss")
	stack.push("he")
	stack.push("asd")
	stack.push("sss")
	stack.push("he")
	stack.push("asd")
	stack.push("sss")
	stack.push("he")
	stack.push("asd")
	stack.push("sss")
	stack.push("he")
	stack.push("asd")
	stack.push("sss")
	stack.push("he")
	stack.push("asd")
	stack.push("sss")
	stack.push("he")
	stack.push("asd")
	stack.push("sss")
	stack.push("he")
	stack.push("asd")
	stack.push("sss")

	t.Logf("%+v", stack)
	t.Log(stack.pop())
	t.Log(stack.pop())
	t.Log(stack.pop())
}
