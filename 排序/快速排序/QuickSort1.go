package QuickSort

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid := depart(arr)
	QuickSort(arr[0:mid])
	QuickSort(arr[mid:])
}

/**
切割 实际就是区分两边的区间
而中间则是最终排完序后的位置
 */
func depart(arr []int) int {
	midValue := arr[len(arr) - 1]

	i, j := 0, 0
	for ; i < len(arr); i ++ {
		if arr[i] < midValue {
			arr[i], arr[j] = arr[j], arr[i]
			j ++
		}
	}
	arr[j], arr[len(arr) - 1] = arr[len(arr) - 1], arr[j]
	return j
}
