package 循环队列

type ArrayQueue struct {
	data []string //队列数据
	head int //队头下标
	tail int //队尾下标
	len int //队列大小
}

func NewArrayQueue(len int) *ArrayQueue {
	return &ArrayQueue{data: make([]string, len), head: 0, tail: 0, len: len}
}

//入队
func (queue *ArrayQueue) enQueue(value string)  {
	//若插入成功, 则队尾的位置变更为(queue.tail + 1) % queue.len, 表示从尾部又循环到头部
	//warning: 当队列满的判断条件是下一个队尾指针和队头指针重合
	if (queue.tail + 1) % queue.len == queue.head {
		return
	}
	//队尾插入
	queue.data[queue.tail] = value
	queue.tail = (queue.tail + 1) % queue.len
}

//出队
func (queue *ArrayQueue) deQueue() string {
	if queue.tail == queue.head {
		return ""
	}
	//获取队头
	ret := queue.data[queue.head]
	//队头向下移动
	queue.head = (queue.head + 1) % queue.len
	return ret
}
