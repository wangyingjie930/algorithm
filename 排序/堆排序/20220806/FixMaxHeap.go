package _0220806

// FixMaxHeap 固定长度的大顶堆, 主要用于动态数据中求topK的场景
type FixMaxHeap struct {
	MaxHeap
	//允许加入的最大长度
	maxLen int
}

func NewFixMaxHeap(arr []int, maxLen int) *FixMaxHeap {
	heap := &FixMaxHeap{maxLen: maxLen}
	heap.data = append([]int{0})
	heap.count = 0

	for _, item := range arr {
		heap.Insert(item)
	}
	return heap
}

func (heap *FixMaxHeap) Insert(item int) {
	if heap.count+1 > heap.maxLen {
		//如果堆已经满了, 并且item比堆顶元素更小, 进行替换
		if heap.data[0] > item {
			heap.data[0] = item
			heap.ShiftDown(0)
		}
	}else {
		//如果堆没有满, 按照正常的堆数据插入
		heap.MaxHeap.Insert(item)
	}
}