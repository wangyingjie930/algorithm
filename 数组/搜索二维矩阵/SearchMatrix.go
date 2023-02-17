/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc: https://leetcode.cn/problems/search-a-2d-matrix/
**/

package 搜索二维矩阵

func searchMatrix(matrix [][]int, target int) bool {
	arrayIndex := binarySearchArray(matrix, target, 0, len(matrix)-1)
	if arrayIndex >= len(matrix) || arrayIndex < 0 {
		return false
	}
	index := binarySearch(matrix[arrayIndex], target, 0, len(matrix[arrayIndex])-1)
	if index >= len(matrix[arrayIndex]) || index < 0 {
		return false
	}
	return matrix[arrayIndex][index] == target
}

// binarySearchArray
//  @Description: 二分查找二维数组
//  @param matrix
//  @param target
//  @param left
//  @param right
//  @return int
func binarySearchArray(matrix [][]int, target int, left, right int) int {
	for left < right {
		mid := (left + right) / 2
		if matrix[mid][0] > target {
			//右区间维护的是x > target
			right = mid - 1
		} else if matrix[mid][len(matrix[mid])-1] < target {
			//左区间维护的是 x < target
			left = mid + 1
		} else {
			//上面的判断都不满足 表示target正处于matrix[mid][0]到matrix[mid][len]的范围内, 直接让left和right选中该区间
			right = mid
			left = mid
		}
	}

	return left
}

// binarySearch
//  @Description: 二分查找一维数组
//  @param matrix
//  @param target
//  @param left
//  @param right
//  @return int
func binarySearch(matrix []int, target int, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
