/**
  @author: wangyingjie
  @since: 2023/2/22
  @desc: https://leetcode.cn/problems/trapping-rain-water/
**/

package 接雨水

func trap(height []int) int {
	maxHeight := height[0]
	for _, h := range height {
		if maxHeight < h {
			maxHeight = h
		}
	}
	res := 0
	//从高层逐层减一, 每一层在左右边界中数格子
	for h := maxHeight; h > 0; h-- {
		//拿到左右边界
		left, right := getBorder(height, h)
		if left >= right {
			continue
		}
		//数一数第h层空出的格子
		for i := left; i <= right; i++ {
			if height[i] < h {
				res++
			}
		}
	}
	return res
}

// getBorder
//  @Description: 获取高度为h的左右边界
//  @param height
//  @param h
//  @return int
//  @return int
func getBorder(height []int, h int) (int, int) {
	left := -1
	right := -1
	for i := 0; i < len(height); i++ {
		if height[i] < h {
			continue
		}
		if left == -1 {
			left = i
		}
		right = i
	}
	return left, right
}
