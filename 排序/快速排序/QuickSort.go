package QuickSort

import (
	InsertSort "algorithm/排序/插入排序"
	"math/rand"
)

type Normal struct {
}

func (n *Normal) Sort(arr []int, start, end int) {
	/*if start >= end {
		return
	}*/
	//优化点2: 当数量足够小时, 启用插入排序
	if end-start < 3 {
		InsertSort.InsertSort(arr[start : end+1])
		return
	}
	p := n.partition(arr, start, end)
	n.Sort(arr, start, p-1)
	n.Sort(arr, p+1, end)
}

func (n *Normal) partition(arr []int, start, end int) int {
	value := arr[start]

	j := start
	k := start + 1
	for ; k <= end; k++ {
		if arr[k] < value {
			arr[j+1], arr[k] = arr[k], arr[j+1]
			j++
		}
	}
	arr[j], arr[start] = arr[start], arr[j]
	return j
}

/*func (n *Normal) partition(arr []int, start, end int) int {
	//优化1: 随机选取一个值作为标定
	//randKey := rand.Int() % (end - start + 1) + start
	//arr[start], arr[randKey] = arr[randKey], arr[start]

	l := start
	j := start
	for i := l + 1; i <= end; i ++ {
		if arr[l] > arr[i] {
			arr[i], arr[j+1] = arr[j+1], arr[i]
			j ++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}*/

func partition2(arr []int, start, end int) int {
	//优化1: 随机选取一个值作为标定
	randKey := rand.Int()%(end-start+1) + start
	arr[start], arr[randKey] = arr[randKey], arr[start]

	l := start
	j := end
	i := start + 1

	//for的多条件使用平行赋值
	for ; true; i, j = i+1, j-1 {
		for i <= end && arr[i] < arr[l] {
			i++
		}
		for j >= start+1 && arr[j] > arr[l] {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[l], arr[j] = arr[j], arr[l]
	return j
}
