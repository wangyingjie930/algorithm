package 合并K个有序链表或数组

import "container/heap"

// Item 节点
type Item struct {
	listIndex int //该元素处于第几个数组
	itemIndex int //属于数组的第几个元素
	value int //元素的值
}

// MaxHeap 堆
type MaxHeap []*Item

func (heap *MaxHeap) Len() int {
	return len(*heap)
}

func (heap *MaxHeap) Less(i, j int) bool {
	return (*heap)[i].value > (*heap)[j].value
}

func (heap MaxHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *MaxHeap) Push(x interface{}) {
	val := x.(*Item)
	*heap = append(*heap, val)
}

func (heap *MaxHeap) Pop() interface{} {
	n := heap.Len()
	ret := (*heap)[n - 1]
	*heap = (*heap)[0:n-1]
	return ret
}

func mergeKArray(lists [][]int) []int {
	h := new(MaxHeap)
	//获取每个数组的第一个数进行堆的构建
	for i := 0; i < len(lists); i++ {
		heap.Push(h, &Item{listIndex: i, itemIndex: 0, value: lists[i][0]})
	}

	var ret []int
	for h.Len() > 0 {
		//弹出堆顶, 放入结果中
		item := heap.Pop(h).(*Item)
		ret = append(ret, item.value)
		//获取弹出节点是从哪个数组过来的, 是第几个元素
		//将所在的数组的下一个元素放入堆中
		if len(lists[item.listIndex]) - 1 > item.itemIndex {
			heap.Push(h, &Item{listIndex: item.listIndex, itemIndex: item.itemIndex + 1, value: lists[item.listIndex][item.itemIndex + 1]})
		}
	}
	return ret
}
