package 循环队列

import (
	"testing"
)

func TestNewArrayQueue(t *testing.T) {
	queue := NewArrayQueue(3)
	queue.enQueue("1")
	t.Log(queue.deQueue())
	queue.enQueue("2")
	t.Log(queue.deQueue())

	queue.enQueue("3")
	queue.enQueue("5")
	queue.enQueue("6")
	queue.enQueue("7")
	t.Log(queue.deQueue())
	t.Log(queue.deQueue())
	t.Log(queue.deQueue())
	t.Log(queue.deQueue())
}
