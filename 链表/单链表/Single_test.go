package 单链表

import "testing"

func TestList_addNode(t *testing.T) {
	list := NewSingleList()
	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)
	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)

	list.deleteKth(2)
	list.Traverse()

	list.removeNthFromEnd(1)
	list.Traverse()
}

func TestList_removeNthFromEnd(t *testing.T) {
	list := NewSingleList()
	list.AddNode(1)
	list.removeNthFromEnd(1)
	list.Traverse()
}