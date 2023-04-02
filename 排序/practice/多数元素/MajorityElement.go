/**
  @author: wangyingjie
  @since: 2023/2/17
  @desc: https://leetcode.cn/problems/majority-element/
**/

package 多数元素

// majorityElement
//  @Description: 方法1: 使用排序, 获取中间位置的即为多数元素
//  @param nums
//  @return int
func majorityElement(nums []int) int {
	heap := newHeap(nums)
	ret := 0
	for i := 0; i <= len(nums)/2; i++ {
		ret = heap.pop()
	}
	return ret
}

type Heap struct {
	arr   []int
	count int
}

func newHeap(nums []int) *Heap {
	arr := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		arr[i+1] = nums[i]
	}
	heap := &Heap{arr: arr, count: len(nums)}
	for i := len(nums) / 2; i > 0; i-- {
		heap.shiftDown(i)
	}
	return heap
}

func (h *Heap) shiftDown(i int) {
	for j := i; j < h.count; {
		minIndex := j
		if j*2 <= h.count && h.arr[2*j] < h.arr[minIndex] {
			minIndex = 2 * j
		}
		if j*2+1 <= h.count && h.arr[2*j+1] < h.arr[minIndex] {
			minIndex = 2*j + 1
		}
		if minIndex == j {
			break
		}
		h.arr[minIndex], h.arr[j] = h.arr[j], h.arr[minIndex]
		j = minIndex
	}
}

func (h *Heap) pop() int {
	tmp := h.arr[1]
	h.arr[1], h.arr[h.count] = h.arr[h.count], h.arr[1]
	h.count--
	h.shiftDown(1)
	return tmp
}
