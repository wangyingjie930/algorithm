/**
  @author: wangyingjie
  @since: 2022/11/27
  @desc: //TODO
**/

package 桶排序

import (
	QuickSort "algorithm/排序/快速排序"
)

func bucketSort(arr []int, bucketNum int) {
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

	//每个桶所表示的区间, 这里减1是怕越界
	interval := (max - min) / (bucketNum - 1)
	buckets := make([][]int, bucketNum)
	for _, v := range arr {
		//将值放入对应的桶中
		buckets[(v - min) / interval] = append(buckets[(v - min) / interval], v)
	}

	var newArr []int
	//每个桶进行快速排序
	for _, bucket := range buckets {
		QuickSort.QuickSort(bucket)
		for _, v := range bucket {
			newArr = append(newArr, v)
		}
	}
	copy(arr, newArr)
}

