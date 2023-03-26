/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc: https://leetcode.cn/problems/search-a-2d-matrix-ii/description/
**/

package 搜索二维矩阵II

// searchMatrix
//  @Description: 关键点就是从右上, 往左或往下走去查找目标值
//  @param matrix
//  @param target
//  @return bool
func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if target == matrix[i][j] {
			return true
		} else if target > matrix[i][j] {
			i++
		} else if target < matrix[i][j] {
			j--
		}
	}
	return false
}
