package _0220806

type MaxHeap struct {
	data []int
	count int
}

func NewMaxHeap(arr []int) MaxHeap {
	heap := MaxHeap{data: []int{0}, count: 0}
	heap.data = append(heap.data, arr...)
	heap.count = len(arr)

	for i := len(arr) / 2; i > 0; i -- {
		heap.ShiftDown(i)
	}
	return heap
}

func (heap *MaxHeap) IsEmpty() bool {
	return heap.count == 0
}

func (heap *MaxHeap) ExtractMax() int {
	max := heap.data[1]
	heap.data[1], heap.data[heap.count] = heap.data[heap.count], heap.data[1]
	heap.count--
	heap.ShiftDown(1)
	return max
}

func (heap *MaxHeap) Insert(i int)  {
	heap.data = append(heap.data, i)
	heap.count++
	heap.ShiftUp(heap.count)
}

func (heap *MaxHeap) ShiftUp(i int) {
	for ; i > 1; i = i / 2 {
		if heap.data[i] > heap.data[i / 2] {
			heap.data[i], heap.data[i / 2] = heap.data[i / 2], heap.data[i]
		}
	}
}

func (heap *MaxHeap) ShiftDown(i int) {
	for j := i; j * 2 <= heap.count; {
		//默认取左节点
		maxIndex := 2 * j
		//左 < 右, 取右
		if j*2+1 <= heap.count && heap.data[j*2] < heap.data[j*2+1] && heap.data[j*2+1] > heap.data[j]{
			maxIndex = j * 2 + 1
		}
		if heap.data[maxIndex] < heap.data[j] {
			break
		}
		//交换
		heap.data[j], heap.data[maxIndex] = heap.data[maxIndex], heap.data[j]
		j = maxIndex
	}
}

// Sort 进行堆排序/**
func (heap *MaxHeap) Sort() []int {
	for ; heap.count > 0; {
		heap.ExtractMax()
	}
	return heap.data[1:]
}