package 中位数

import "container/heap"

type MaxHeap []int
type MinHeap struct {
	MaxHeap
}

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {
	val := x.(int)
	*m = append(*m, val)
}

func (m *MaxHeap) Pop() interface{} {
	n := m.Len()
	ret := (*m)[n - 1]
	*m = (*m)[0:n-1]
	return ret
}

func (m MinHeap) Less(i, j int) bool {
	return m.MaxHeap[i] < m.MaxHeap[j]
}


type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
	count int
}


func Constructor() MedianFinder {
	return MedianFinder{minHeap: &MinHeap{}, maxHeap: &MaxHeap{}, count: 0}
}

// AddNum 添加数据
func (this *MedianFinder) AddNum(num int)  {
	if this.count == 0 {
		heap.Push(this.maxHeap, num)
		this.count ++
		return
	}

	topMax := (*this.maxHeap)[0]
	//1. 往哪里放
	if num <= topMax {
		//小于大顶堆的往大顶堆放
		heap.Push(this.maxHeap, num)
	}else {
		//其他的往小顶堆放
		heap.Push(this.minHeap, num)
	}
	this.count++

	//2. 调整堆, 维持定义, 实际就是维持两个区间, 一个是小于中位数的区间, 另一个是大于中位数的区间
	if this.maxHeap.Len() - this.minHeap.Len() != 0 && this.count % 2 == 0 {
		//偶数个的时候, 两个堆的数量必须相等
		val := heap.Pop(this.maxHeap).(int)
		heap.Push(this.minHeap, val)
	}else if this.maxHeap.Len() - this.minHeap.Len() != 1 && this.count % 2 == 1 {
		//奇数个的时候, 大顶堆一定要比小顶堆多出一个
		val := heap.Pop(this.minHeap).(int)
		heap.Push(this.maxHeap, val)
	}
}

// FindMedian  获取中位数
func (this *MedianFinder) FindMedian() float64 {
	if this.count%2 == 1 {
		//注意不要使用pop
		val := (*this.maxHeap)[0]
		return float64(val)
	}else {
		val1 := (*this.maxHeap)[0]
		val2 := this.minHeap.MaxHeap[0]
		return (float64(val2) + float64(val1)) / 2
	}
}
