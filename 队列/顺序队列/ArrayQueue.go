package 顺序队列

type ArrayQueue struct {
	data []string //队列数据
	head int //队头下标
	tail int //队尾下标, 指向的是data最后元素下标+1
	len int //队列大小
}

func NewArrayQueue(len int) *ArrayQueue {
	return &ArrayQueue{data: make([]string, len), head: 0, tail: 0, len: len}
}

//入队
func (queue *ArrayQueue) enQueue(value string)  {
	if queue.tail - queue.head == queue.len {
		return
	}
	//当队尾节点已经到达队尾, 并且队头的下标不为0, 表示前面还有剩余的位置可以插入
	//进行队列的搬移
	if queue.tail == queue.len && queue.head > 0 {
		queue.migrate()
	}
	//队尾插入
	queue.data[queue.tail] = value
	//队尾下移
	queue.tail++
}

//需要进行数据迁移
func (queue *ArrayQueue) migrate() {
	newData := make([]string, queue.len)
	//将队头和队尾之间的内容搬移
	copy(newData, queue.data[queue.head:queue.tail])
	queue.data = newData
	//队尾指针移动
	queue.tail = queue.tail - queue.head
	//队头移动
	queue.head = 0
}

//出队
func (queue *ArrayQueue) deQueue() string {
	if queue.tail == queue.head {
		return ""
	}
	//获取队头
	ret := queue.data[queue.head]
	//队头向下移动
	queue.head ++
	return ret
}


