package InsertSort

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		tmp := arr[i]

		for ; j >= 0 && arr[j] > tmp; j-- {
			arr[j+1] = arr[j]
		}
		
		arr[j+1] = tmp
	}
}
