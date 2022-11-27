package SelectSort

func SelectionSort(arr []int) []int {
	for k, v := range arr {
		minIndex := k
		min := v
		for j := k + 1; j < len(arr); j ++ {
			if arr[j] < min {
				minIndex = j
				min = arr[j]
			}
		}
		arr[k], arr[minIndex] = arr[minIndex], arr[k]
	}
	return arr
}
