package 计数排序

func CountingSort(arr []int)  {
	if len(arr) <= 1 {
		return
	}


	//准备: 获得最大最小值
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if min > v {
			min = v
		}
		if max < v{
			 max = v
		}
	}

	//准备: 计数数组的统计
	countingArr := make([]int, max - min + 1)
	for _, v := range arr {
		//统计每个数出现的次数
		countingArr[v - min] ++
	}
	//准备: 累加数组的统计
	for i := 1; i < len(countingArr); i ++ {
		//进行一次累加次数, 用于稳定的排序
		//相当于给出了arr中在排序之后每个元素最后出现的排序位置
		countingArr[i] = countingArr[i] + countingArr[i - 1]
	}

	//进行计数排序
	newArr := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		//从后往前遍历, 配合累加数组
		//countingArr[arr[i] - min] - 1表示的arr[i]在排序后应该出现在哪个位置, 数组以0开头, 需要-1
		newArr[countingArr[arr[i] - min] - 1] = arr[i]
		//自减1, 相当于最后出现的位置往前移了一些
		countingArr[arr[i] - min]--
	}
	copy(arr, newArr)
}