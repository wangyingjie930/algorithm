/**
  @author: wangyingjie
  @since: 2023/3/21
  @desc: https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
  掌握程度: 需要重写
**/

package 寻找两个正序数组的中位数

// findMedianSortedArrays
//  @Description: 将取合并后的中位数转化为求第k((n1+n2)/2)个问题
//  @param nums1
//  @param nums2
//  @return float64
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		//为奇数时, 取第 n / 2 + 1个
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		//为偶数时, 取第n/2, n/2+1个取平均值
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

// getKthElement
//  @Description: 获取两个有序数组合并后的第k个, 注意k表示为第k个, 然后下标的转化要-1, 一些二分要 k/2
//  @param nums1
//  @param nums2
//  @param k
//  @return int
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			//当某个数组比较完成是, 剩下的第k个即为合并之后的第k个
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			//合并之后的第1个可以直接比较两个开头的数组即可
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2                                 //比较两个数组的前k/2个
		newIndex1 := min(index1+half, len(nums1)) - 1 //转成下标
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1 // 将k减去不符合的个数
			index1 = newIndex1 + 1      //从新的下标开始
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
