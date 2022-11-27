package practice

/**
寻找两个正序数组的中位数
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

类似于归并算法中的merge函数
 */
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	index := length / 2
	arr := make([]int, length)
	i := 0
	j := 0
	tmp := 0
	for i < len(nums1) || j < len(nums2) {
		if j >= len(nums2) {
			arr[tmp] = nums1[i]
			i ++
		}else if i >= len(nums1) {
			arr[tmp] = nums2[j]
			j ++
		}else if nums1[i] <= nums2[j] {
			arr[tmp] = nums1[i]
			i ++
		}else if nums1[i] > nums2[j] {
			arr[tmp] = nums2[j]
			j ++
		}
		if tmp == index && length % 2 == 0 {
			return (float64(arr[index - 1]) + float64(arr[index])) / 2
		}else if tmp == index && length % 2 == 1 {
			return float64(arr[tmp])
		}
		tmp ++
	}
	return -1
}

/**
维护大根堆和小根堆
1. 第一个数进大根堆
2. 当前数 <= 大根堆堆顶元素, 入大根堆, 否则入小根堆
3. 当任意堆长度差超过2, 从长的堆将堆顶元素调到另外的堆上
 */