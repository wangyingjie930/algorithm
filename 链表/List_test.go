package List

import (
	"fmt"
	"testing"
)

var arr = []int{92,4,2,49,47,32,90,95,35,85,41,84,82,48,6,94,69,20,50,87,79,28,
	67,15,36,55,75,93,91,56,8,12,42,16,66,59,57,5,33,1,54,97,80,77,65,52,29,81,14,13,
	86,64,98,68,46,58,76,70,11,96,60,31,34,44,30,22,61,24,43,26,62,21,74,83,39,72,73,88,45,27,0,10,89,
	51,18,37,17,3,40,9,19,63,71,38,53,78,7,23,99,25}
var head = NewList(arr)

/**
创建链表测试用例
 */
func TestNewList(t *testing.T) {
	head := NewList(arr)
	ShowList(head)
	fmt.Printf("%+v", head)
}

func TestReverseList(t *testing.T) {
	ShowList(head)
	head = ReverseList(head)
	ShowList(head)
}

func TestDetectCycle(t *testing.T) {
	node1 := &Node{Val: 3}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 0}
	node4 := &Node{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node1
	fmt.Println(DetectCycle(node1))
}

func TestGetIntersectionNode(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 9}
	node3 := &Node{Val: 1}
	node4 := &Node{Val: 2}
	node5 := &Node{Val: 4}
	node6 := &Node{Val: 3}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node6.Next = node4
	fmt.Println(GetIntersectionNode(node1, node6))
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *Node
		l2 *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "test1", args: args{l1: NewList([]int{1,3,5,7}), l2: NewList([]int{2, 4, 6, 8})}, want: NewList([]int{1,2,3,4,5,6,7,8})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.l1, tt.args.l2)
			t.Logf("mergeTwoLists() = %v", got)
		})
	}
}

func Test_middleNode(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "test1", args: args{head: NewList([]int{1,3,5})}, want: NewList([]int{3})},
		{name: "test2", args: args{head: NewList([]int{1,3,6,5,4,7})}, want: NewList([]int{3})},
	}
	for _, tt := range tests {
		got := middleNode(tt.args.head)
		t.Logf("middleNode() = %v", got)
	}
}