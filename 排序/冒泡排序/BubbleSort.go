package 冒泡排序

func Sort(arr []int) {
	start := 0
	end := len(arr) - 1

	for j := start; j <= end; j ++{
		flag := false
		for i := start; i <= end - j - 1; i ++ {
			if arr[i] > arr[i + 1] {
				arr[i], arr[i + 1] = arr[i + 1], arr[i]
				flag = true
			}
		}

		if !flag {
			return
		}
	}
}
