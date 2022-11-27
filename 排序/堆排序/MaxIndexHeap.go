package Heap

import (
	"fmt"
)

type MaxIndexHeap struct {
	/**
	数据保持不变
	 */
	data    []int
	count   int
	/**
	* 索引堆, 对索引建堆
	* 堆键->data键->data值
	 */
	indexes []int
	reverse []int
}

func NewMaxIndexHeap () MaxIndexHeap {
	data := []int{0}
	index := []int{0}
	reverse := []int{0}
	return MaxIndexHeap{data: data, count: 0, indexes: index, reverse: reverse}
}

func (h *MaxIndexHeap) Size() int  {
	return h.count
}

func (h *MaxIndexHeap) IsEmpty() bool {
	return h.count == 0
}

func (h *MaxIndexHeap) Data()  {
	fmt.Println("data: ", h.data)
	fmt.Println("index: ", h.indexes)
}

/**
大顶堆插入数据
 */
func (h *MaxIndexHeap) Insert(item int)  {
	//h.data[h.count + 1] = item
	h.data = append(h.data, item)
	h.indexes = append(h.indexes, h.count)

	h.count ++
	h.shiftUp(h.count)
}

/**
插入完之后重新排列, 使它符合大顶堆条件
 */
func (h *MaxIndexHeap) shiftUp(i int) {
	for ;i > 1 && h.data[h.indexes[i]] > h.data[h.indexes[ i / 2]]; i = i / 2 {
		h.indexes[i], h.indexes[ i / 2] = h.indexes[i / 2], h.indexes[i]
	}
}

func (h *MaxIndexHeap) ExtractMax() int  {
	max := h.indexes[1]
	h.indexes[1], h.indexes[h.count] = h.indexes[h.count], h.indexes[1]
	h.count --
	h.shiftDown(1)
	return h.data[max]
}

func (h *MaxIndexHeap) ExtractMaxIndex() int  {
	max := h.indexes[1]
	h.indexes[1], h.indexes[h.count] = h.indexes[h.count], h.indexes[1]
	h.count --
	h.shiftDown(1)
	return max
}

func (h *MaxIndexHeap) shiftDown(i int)  {
	for ;2 * i <= h.count; {
		//默认取左节点
		j := 2 * i
		//左 < 右, 取右
		if j + 1 <= h.count && h.data[h.indexes[j]] < h.data[h.indexes[j + 1]] {
			j = j + 1
		}
		//当前节点大于左右节点, 跳出
		if h.data[h.indexes[i]] >= h.data[h.indexes[j]] {
			break
		}
		//当前节点与 左/右 节点交换
		h.indexes[i], h.indexes[j] = h.indexes[j], h.indexes[i]
		i = j //向前移动
	}
}

func (h *MaxIndexHeap) GetItem (i int) int {
	return h.data[i]
}

func (h *MaxIndexHeap) Change(i int, item int) {
	//这个i是data键
	i = i + 1
	h.data[i] = item
	var j int

	//查找index堆键
	for k, v := range h.indexes {
		if v == i {
			j = k
			break
		}
	}
	h.shiftDown(j)
	h.shiftUp(j)
}

