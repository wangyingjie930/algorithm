package _0220806

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func NewPriorityQueue(items map[string]int) *PriorityQueue {
	var pq PriorityQueue
	i := 0
	for value, priority := range items {
		pq = append(pq, &Item{value: value, priority: priority, index: i})
		i ++
	}
	heap.Init(&pq)
	return &pq
}

func (pq *PriorityQueue) Println ()  {
	for _, v := range *pq {
		fmt.Printf("%+v\n", v)
	}
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func mergeKSortList(lists []map[string]int) []string {
	//var pq *PriorityQueue
	pq := &PriorityQueue{}
	i := 0
	for _, list := range lists {
		for value, priority := range list {
			heap.Push(pq, &Item{value: value, priority: priority, index: i})
		}
	}

	var ret []string
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		ret = append(ret, item.value)
	}
	return ret
}
