package InsertSort

func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i ++ {
		for j := i; j > 0; j-- {
			if arr[j - 1] >= arr[j] {
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
			}  else {
				break
			}
		}
	}
	return arr
}