/**
  @author: wangyingjie
  @since: 2023/4/3
  @desc: 重写练习
**/

package _0230403

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if (len1+len2)%2 == 1 {
		return float64(getKthElement(nums1, nums2, (len1+len2)/2+1))
	} else {
		e1 := getKthElement(nums1, nums2, (len1+len2)/2+1)
		e2 := getKthElement(nums1, nums2, (len1+len2)/2)
		return float64(e1+e2) / float64(2)
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	//记住四个变量维护两个区间
	l1, l2 := 0, 0
	for {
		if l1 == len(nums1) {
			return nums2[l2+k-1]
		}
		if l2 == len(nums2) {
			return nums1[l1+k-1]
		}
		if k == 1 {
			return min(nums1[l1], nums2[l2])
		}

		half := k / 2
		r1, r2 := min(l1+half, len(nums1))-1, min(l2+half, len(nums2))-1
		if nums1[r1] < nums2[r2] {
			k = k - (r1 - l1 + 1)
			l1 = r1 + 1
		} else {
			k = k - (r2 - l2 + 1)
			l2 = r2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
