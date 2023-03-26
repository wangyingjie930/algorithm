/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc: https://leetcode.cn/problems/rotate-image/
**/

package 旋转图像

func rotate(matrix [][]int) {
	//沿左上-右下的轴进行对称翻转
	for i := 0; i < len(matrix); i++ {
		for j := i; j >= 0; j-- {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	//沿中线翻转
	for i := 0; i < len(matrix); i++ {
		for left, right := 0, len(matrix[i])-1; left <= right; left, right = left+1, right-1 {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
		}
	}
}
