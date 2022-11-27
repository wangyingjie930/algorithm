package QuickSort

import (
	InsertSort "imooc-product/backend/排序/插入排序"
	"math/rand"
)

type ThreeWays struct {}

func (t *ThreeWays) Sort(arr []int, start, end int) {
	//优化点2: 当数量足够小时, 启用插入排序
	if end - start < 3 {
		InsertSort.InsertSort(arr[start:end + 1])
		return
	}
	lt, gt := t.partition(arr, start, end)
	t.Sort(arr, start, lt - 1)
	t.Sort(arr, gt, end)
}

func (t *ThreeWays) partition(arr []int, l, r int) (int, int) {
	randKey := rand.Int() % (r - l + 1) + l
	arr[l], arr[randKey] = arr[randKey], arr[l]

	lt := l
	gt := r + 1
	i := l + 1
	for i < gt {
		if arr[i] < arr[l] {
			arr[i], arr[lt + 1] = arr[lt + 1], arr[i]
			lt ++
			i ++
		}else if arr[i] > arr[l] {
			arr[i], arr[gt - 1] = arr[gt - 1], arr[i]
			gt --
		}else {
			i ++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	return lt, gt
}
