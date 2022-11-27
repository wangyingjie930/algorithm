package MergeSort
/**
20221126 自测
 */
func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	mergeSort(arr[0:mid])
	mergeSort(arr[mid:])

	copy(arr, merge1(arr[0:mid], arr[mid:]))
}

/**
合并两个有序数组
 */
func merge1(arr1, arr2 []int) []int {
	newArr := make([]int, len(arr1) + len(arr2))
	i, j, k := 0, 0, 0
	for ; i < len(arr1) && j < len(arr2); k ++{
		if arr1[i] < arr2[j] {
			newArr[k] = arr1[i]
			i++
		}else if arr1[i] > arr2[j] {
			newArr[k] = arr2[j]
			j ++
		}else {
			i ++
			j ++
		}
	}

	for ; i >= len(arr1) && j < len(arr2); j, k = j + 1, k + 1 {
		newArr[k] = arr2[j]
	}

	for ; j >= len(arr2) &&  i < len(arr1); i, k = i + 1, k + 1 {
		newArr[k] = arr1[i]
	}
	return newArr
}
