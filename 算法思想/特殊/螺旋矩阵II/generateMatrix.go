/**
  @author: wangyingjie
  @since: 2023/3/4
  @desc: https://leetcode.cn/problems/spiral-matrix-ii/
**/

package 螺旋矩阵II

// generateMatrix
//  @Description: 过程正好相反
//  @param n
//  @return [][]int
func generateMatrix(n int) [][]int {
	left, right, upper, down := 0, n-1, 0, n-1
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	curNum := 1

	for true {
		//左上->右上
		for i := left; i <= right; i++ {
			matrix[upper][i] = curNum
			curNum++
		}
		upper++ //上边界已经遍历, 向下缩
		if upper > down {
			break
		}

		//右上->右下
		for i := upper; i <= down; i++ {
			matrix[i][right] = curNum
			curNum++
		}
		right-- //右边界已经遍历, 向左缩一格
		if right < left {
			break
		}

		//右下->左下
		for i := right; i >= left; i-- {
			matrix[down][i] = curNum
			curNum++
		}
		down-- //下边界已经遍历, 向上缩
		if down < upper {
			break
		}

		//左下->左上
		for i := down; i >= upper; i-- {
			matrix[i][left] = curNum
			curNum++
		}
		left++
		if left > right {
			break
		}
	}

	return matrix
}
