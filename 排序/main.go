package main

import (
	"math/rand"
	"time"
)

func main() {
	//Provide seed
	rand.Seed(time.Now().Unix())

	//Generate a random array of length n
	//arr := rand.Perm(20)
	//MergeSort.MergeSort(arr, 0, len(arr) - 1)
	//QuickSort.QuickSort(arr, 0, len(arr) - 1)
	/*new(QuickSort.ThreeWays).Sort(arr, 0, len(arr) - 1)
	fmt.Println(arr)

	arr = rand.Perm(20)
	new(QuickSort.Normal).Sort(arr, 0, len(arr) - 1)
	fmt.Println(arr)*/

/*	heap := Heap.NewMaxHeap()
	for _, v := range rand.Perm(50) {
		heap.Insert(v)
	}
	heap.Data()

	heap = Heap.NewMaxHeapHeapify(rand.Perm(50))
	heap.Data()

	for !heap.IsEmpty() {
		fmt.Print(heap.ExtractMax(), " ")
	}

	arr := rand.Perm(50)
	Heap.HeapSort(arr, 50)
	fmt.Print("\n", arr)*/

	/*heap := Heap.NewMaxIndexHeap()
	for _, v := range rand.Perm(50) {
		heap.Insert(v)
	}

	heap.Change(4, 100)
	for !heap.IsEmpty() {
		fmt.Print(heap.ExtractMax(), " ")
	}*/

	/*topN := Heap.NewTopN(1000)
	for _, v := range rand.Perm(100000000) {
		topN.Insert(v)
	}
	topN.Data()*/
}
