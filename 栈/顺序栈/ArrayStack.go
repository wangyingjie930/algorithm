package 顺序栈

type ArrayStack struct {
	data []string
	count int //个数
	len int //容量
}

func NewArrayStack(len int) *ArrayStack {
	return &ArrayStack{data: make([]string, len), len: len, count: 0}
}

//入栈
func (stack *ArrayStack) push(value string) {
	if stack.count == stack.len {
		//动态扩容
		stack.resize()
	}
	stack.data[stack.count] = value
	stack.count ++
}

func (stack *ArrayStack) resize() {
	newData := make([]string, stack.len * 2)
	copy(newData, stack.data)
	stack.data = newData
	stack.len = stack.len * 2
}

//出栈
func (stack *ArrayStack) pop() string {
	if stack.count == 0 {
		return ""
	}
	ret := stack.data[stack.count - 1]
	stack.data[stack.count - 1] = ""
	stack.count --
	return ret
}



