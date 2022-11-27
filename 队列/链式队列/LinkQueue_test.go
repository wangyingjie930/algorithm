package 链式队列

import (
	"testing"
)

func TestNewLinkQueue(t *testing.T) {
	queue := NewLinkQueue()
	queue.enQueue("1")
	queue.enQueue("2")
	queue.enQueue("3")
	queue.enQueue("2")
	queue.enQueue("3")
	queue.enQueue("2")
	queue.enQueue("3")
	queue.enQueue("2")
	queue.enQueue("3")
	queue.enQueue("2")
	queue.enQueue("3")

	t.Log(queue.deQueue(), queue.deQueue(), queue.deQueue())
}
