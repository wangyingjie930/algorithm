package Heap

import "fmt"

type MaxHeap struct {
	data []int
	count int
}

func NewMaxHeap () MaxHeap {
	data := []int{0}
	return MaxHeap{data: data, count: 0}
}

/**
 Heapify构建
 */
func NewMaxHeapHeapify(arr []int) MaxHeap  {
	data := []int{0}
	data = append(data, arr...)
	count := len(arr)
	heap := MaxHeap{data: data, count: count}

	//count / 2 表示最后一个非叶子结点的索引位置
	//依次向上减, 依次进行shiftDown操作
	for i := count / 2; i > 0; i -- {
		heap.shiftDown(i)
	}
	return heap
}

func (h *MaxHeap) Size() int  {
	return h.count
}

func (h *MaxHeap) IsEmpty() bool {
	return h.count == 0
}

func (h *MaxHeap) Data()  {
	fmt.Println(h.data)
}

/**
大顶堆插入数据
 */
func (h *MaxHeap) Insert(item int)  {
	//h.data[h.count + 1] = item
	h.data = append(h.data, item)
	h.count ++
	h.shiftUp(h.count)
}

/**
插入完之后重新排列, 使它符合大顶堆条件
 */
func (h *MaxHeap) shiftUp(i int) {
	for ;i > 1 && h.data[i] > h.data[i / 2]; i = i / 2 {
		h.data[i], h.data[i / 2] = h.data[i / 2], h.data[i]
	}
}

func (h *MaxHeap) ExtractMax() int  {
	max := h.data[1]
	h.data[1], h.data[h.count] = h.data[h.count], h.data[1]
	h.count --
	h.shiftDown(1)
	return max
}

func (h *MaxHeap) shiftDown(i int)  {
	for ;2 * i <= h.count; {
		//默认取左节点
		j := 2 * i
		//左 < 右, 取右
		if j + 1 <= h.count && h.data[j] < h.data[j + 1] {
			j = j + 1
		}
		//当前节点大于左右节点, 跳出
		if h.data[i] >= h.data[j] {
			break
		}
		//当前节点与 左/右 节点交换
		h.data[i], h.data[j] = h.data[j], h.data[i]
		i = j //向前移动
	}
}

