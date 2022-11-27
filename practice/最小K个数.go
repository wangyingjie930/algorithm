package practice

import "math/rand"

/**
使用快速排序实现
 */
func SmallestK(arr []int, k int) []int {
	sort(arr, 0, len(arr) - 1)
	return arr
}

func sort(arr []int, start, end int) {
	if end < 0 || start < 0 || start >= end {
		return
	}
	if end - start == 1 {
		if arr[start] > arr[end] {
			arr[start], arr[end] = arr[end], arr[start]
		}
		return
	}
	lt, gt := partition(arr, start, end)
	sort(arr, start, lt)
	sort(arr, gt, end)
}

func partition(arr []int, start, end int) (int, int) {
	/*v := arr[start]
	j := start //表示比v大的最后临界点, j+1之后的都比v大
	for i := start + 1; i <= end; i++ {
		if arr[i] < v {
			arr[j+1], arr[i] = arr[i], arr[j+1]
			j ++
		}
	}
	arr[start], arr[j] = arr[j], arr[start]
	return j*/

	//优化:
	randKey := rand.Int() % (end - start + 1) + start
	arr[start], arr[randKey] = arr[randKey], arr[start]

	lt := start
	gt := end + 1
	i := start + 1
	for i < gt {
		if arr[i] < arr[start] {
			arr[i], arr[lt + 1] = arr[lt + 1], arr[i]
			lt ++
			i ++
		}else if arr[i] > arr[start] {
			arr[i], arr[gt - 1] = arr[gt - 1], arr[i]
			gt --
		}else {
			i ++
		}
	}
	arr[start], arr[lt] = arr[lt], arr[start]
	return lt, gt
}