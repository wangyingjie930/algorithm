/**
  @author: wangyingjie
  @since: 2023/3/4
  @desc: https://leetcode.cn/problems/spiral-matrix/description/
  https://leetcode.cn/problems/spiral-matrix/solutions/7155/cxiang-xi-ti-jie-by-youlookdeliciousc-3/
**/

package 螺旋矩阵

// spiralOrder
//  @Description: 明确好上下左右边界遍历
//  @param matrix
//  @return []int
func spiralOrder(matrix [][]int) []int {
	left, right, upper, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	var ret []int
	for true {
		//左上->右上
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[upper][i])
		}
		upper++ //上边界已经遍历, 向下缩
		if upper > down {
			break
		}

		//右上->右下
		for i := upper; i <= down; i++ {
			ret = append(ret, matrix[i][right])
		}
		right-- //右边界已经遍历, 向左缩一格
		if right < left {
			break
		}

		//右下->左下
		for i := right; i >= left; i-- {
			ret = append(ret, matrix[down][i])
		}
		down-- //下边界已经遍历, 向上缩
		if down < upper {
			break
		}

		//左下->左上
		for i := down; i >= upper; i-- {
			ret = append(ret, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return ret
}
