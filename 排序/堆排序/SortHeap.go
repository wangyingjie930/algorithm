package Heap

import "fmt"

/**
数组原地进行堆排序
 */
func HeapSort (arr []int, n int)  {
	//构建大顶堆
	fmt.Println("origin:", arr)
	for i := (n - 1) / 2; i >= 0; i --  {
		shiftDown(arr, n, i)
	}
	fmt.Println("heap:", arr)
	for i := n - 1; i > 0; i -- {
		arr[0], arr[i] = arr[i], arr[0]
		shiftDown(arr, i, 0)
	}
	fmt.Println(arr)
}

func shiftDown (arr []int, n, i int) {
	for i < (n - 1) / 2 {
		l := 2 * i + 1
		r := 2 * i + 2
		k := l

		if r < n && arr[k] < arr[r] {
			k = r
		}
		if arr[i] >= arr[k] {
			break
		}
		arr[i], arr[k] = arr[k], arr[i]
		i = k
	}
}
