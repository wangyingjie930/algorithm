package 基数排序

import "math"

// RadixSort
//  @Description:
//  @param arr
func RadixSort(arr []int) {
	//准备: 获得最大最小值
	max := arr[0]
	for _, v := range arr {
		if max < v{
			max = v
		}
	}

	//统计位数
	width := 0
	for i := max; i > 0; i = i / 10 {
		width++
	}

	for i := 1; i <= width; i++ {
		//使用计数排序
		countingArr := make([]int, 10)
		//  计数数组
		for _, v := range arr {
			//获取第i位的数字
			index := getDigit(v, i)
			countingArr[index]++
		}
		//  累加数组, 必须保证稳定性基数排序才能生效!!
		for j := 1; j < len(countingArr); j++ {
			countingArr[j] = countingArr[j] + countingArr[j - 1]
		}
		//  对第N位的数字进行计数排序
		newArr := make([]int, len(arr))
		for k := len(arr) - 1; k >= 0; k -- {
			newArr[countingArr[getDigit(arr[k], i)] - 1] = arr[k]
			countingArr[getDigit(arr[k], i)]--
		}
		copy(arr, newArr)
	}
}

// getDigit
//  @Description: 获取数字个位,十位,百位的数字
//  @param num 被操作的数字
//  @param pos 第几位
//  @return int
func getDigit(num int, pos int) int {
	// 10的pos次方
	v := int(math.Pow(10, float64(pos)))
	return (num % v) / (v / 10)
}
