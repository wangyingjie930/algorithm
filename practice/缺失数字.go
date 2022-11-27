package practice

import (
	"container/heap"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
并不是最优解, 写着玩的

1. 将数组构建成青哥大根堆
2. 弹出堆顶元素,
    如果第一个堆顶元素不是n, 则返回n
    如果弹出的元素不是连续的, 则获得这个数
 */
func MissingNumber(nums []int) int {
	h := IntHeap(nums)
	heap.Init(&h)
	i := 1
	tmp := math.MinInt32
	for h.Len() > 0 {
		val := heap.Pop(&h).(int)
		if i == 1 && val != len(nums) {
			return len(nums)
		}
		if i != 1 && tmp - val != 1 {
			return tmp - 1
		}
		i ++
		tmp = val
	}
	return 0
}
