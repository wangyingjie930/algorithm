package MergeSort

func sortArray(nums []int) []int {
	helper(nums, 0, len(nums)-1)
	return nums
}

func helper(nums []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	helper(nums, start, mid)
	helper(nums, mid+1, end)
	newMerge(nums, start, mid, end)
}

//func merge(nums []int, start, mid, end int) {
//	var tmp []int //先复制数据出来, 因为后面要修改的是nums的值
//	for i := start; i <= end; i++ {
//		tmp = append(tmp, nums[i])
//	}
//
//	//维护的是新复制出来的指针关系
//	i1 := 0
//	i2 := mid - start + 1
//	newMid := mid - start
//
//	for i := start; i <= end; i++ {
//		if i1 > newMid {
//			nums[i] = tmp[i2]
//			i2++
//		} else if i2 >= len(tmp) {
//			nums[i] = tmp[i1]
//			i1++
//		} else if tmp[i1] > tmp[i2] {
//			nums[i] = tmp[i2]
//			i2++
//		} else if tmp[i1] <= tmp[i2] {
//			nums[i] = tmp[i1]
//			i1++
//		}
//	}
//}

func newMerge(nums []int, start, mid, end int) {
	tmpNums := make([]int, end-start+1)
	copy(tmpNums, nums[start:end+1])

	i := 0
	j := mid + 1 - start
	k := start
	for ; i <= mid-start && j <= end-start; k++ {
		if tmpNums[i] < tmpNums[j] {
			nums[k] = tmpNums[i]
			i++
		} else {
			nums[k] = tmpNums[j]
			j++
		}
	}

	for i <= mid-start {
		nums[k] = tmpNums[i]
		i, k = i+1, k+1
	}

	for j <= end-start {
		nums[k] = tmpNums[j]
		j, k = j+1, k+1
	}
}
