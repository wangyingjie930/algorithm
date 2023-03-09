/**
  @author: wangyingjie
  @since: 2023/3/10
  @desc: https://leetcode.cn/problems/merge-k-sorted-lists/description/
**/

package 合并K个升序链表

import List "algorithm/链表/单链表"

func mergeKLists(lists []*List.Node) *List.Node {
	heap := newMinHeap()

	dummy := &List.Node{Val: -1, Next: nil}
	pre := dummy

	for i := 0; i < len(lists); i++ {
		heap.push(lists[i])
	}
	for len(heap.nums) > 1 {
		tmp := heap.pop()
		if tmp.Next != nil {
			heap.push(tmp.Next)
		}
		pre.Next = tmp
		pre = pre.Next
	}
	return dummy.Next
}

// MinHeap 小顶堆
type minHeap struct {
	nums []*List.Node
}

func newMinHeap() *minHeap {
	return &minHeap{nums: []*List.Node{nil}}
}

func (h *minHeap) push(num *List.Node) {
	if num == nil {
		return
	}
	h.nums = append(h.nums, num)
	h.shiftUp(len(h.nums) - 1)
}

func (h *minHeap) pop() *List.Node {
	if len(h.nums) == 1 {
		return nil
	}
	tmp := h.nums[1]
	h.nums[1] = h.nums[len(h.nums)-1]
	h.nums = h.nums[0 : len(h.nums)-1]
	h.shiftDown(1)
	return tmp
}

// shiftUp
//  @Description: 向上堆化
//  @receiver h
//  @param index
func (h *minHeap) shiftUp(index int) {
	for ; index/2 >= 1; index = index / 2 {
		if h.nums[index].Val < h.nums[index/2].Val {
			h.nums[index], h.nums[index/2] = h.nums[index/2], h.nums[index]
		}
	}
}

// shiftDown
//  @Description: 向下堆化
//  @receiver h
//  @param index
func (h *minHeap) shiftDown(index int) {
	for index*2 < len(h.nums) {
		minIndex := index * 2
		if len(h.nums) > index*2+1 && h.nums[2*index+1].Val < h.nums[2*index].Val {
			minIndex = index*2 + 1
		}
		if h.nums[minIndex].Val < h.nums[index].Val {
			h.nums[minIndex], h.nums[index] = h.nums[index], h.nums[minIndex]
		}
		index = minIndex
	}
}
