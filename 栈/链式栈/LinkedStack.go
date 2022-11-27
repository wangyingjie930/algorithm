package 链式栈

type Node struct {
	Val string
	Next *Node
}

type LinkedStack struct {
	top *Node
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{top: &Node{Val: "", Next: nil}}
}

// Push 入栈
func (stack *LinkedStack) Push(value string) {
	stack.top = &Node{Val: value, Next: stack.top}
}

// Pop 出栈
func (stack *LinkedStack) Pop()string {
	if stack.top.Val == "" {
		return ""
	}
	top := stack.top
	stack.top = stack.top.Next
	return top.Val
}

