package MergeSort

func MergeSort(arr []int, start int, end int)  {
	if start >= end {
		return
	}
	//优化点2: 当数量足够小时, 启用插入排序
	/*if end - start < 3 {
		InsertSort.InsertSort(arr[start:end+1])
		return
	}*/

	mid := (start + end) / 2
	MergeSort(arr, start, mid)
	MergeSort(arr, mid + 1, end)

	//这个判断是优化点1, 只要arr[mid] > arr[mid + 1]的时候, 就要进行排序
	if arr[mid] > arr[mid + 1] {
		merge(arr, start, mid, end)
	}
}

/**
自底向上的归并排序
 */
func MergeSortBU(arr []int) {
	//
}

/**
合并两个有序数组
*/
func merge(arr []int, start, mid, end int)  {
	aux := make([]int, end - start + 1)
	for i := start; i <= end; i++ {
		aux[i - start] = arr[i]
	}
	i := start
	j := mid + 1

	for k := start; k <= end; k ++ {
		if i > mid {
			arr[k] = aux[j - start]
			j ++
		}else if j > end {
			arr[k] = aux[i - start]
			i ++
		}else if aux[i - start] < aux[j - start] {
			arr[k] = aux[i - start]
			i ++
		}else {
			arr[k] = aux[j - start]
			j ++
		}
	}
}
