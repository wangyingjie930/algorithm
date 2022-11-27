package 单链表

import "testing"

func TestList_addNode(t *testing.T) {
	list := NewSingleList()
	list.addNode(1)
	list.addNode(2)
	list.addNode(3)
	list.addNode(1)
	list.addNode(2)
	list.addNode(3)

	list.deleteKth(2)
	list.traverse()

	list.removeNthFromEnd(1)
	list.traverse()
}

func TestList_removeNthFromEnd(t *testing.T) {
	list := NewSingleList()
	list.addNode(1)
	list.removeNthFromEnd(1)
	list.traverse()
}