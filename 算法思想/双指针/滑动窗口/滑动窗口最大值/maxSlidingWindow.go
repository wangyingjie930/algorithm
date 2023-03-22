/**
  @author: wangyingjie
  @since: 2023/3/23
  @desc: https://leetcode.cn/problems/sliding-window-maximum/
  掌握程度: 不熟
**/

package 滑动窗口最大值

func maxSlidingWindow(nums []int, k int) []int {
	var rankQueue []int //保持单调队列是递减, 这样能保持当最大值弹出时可以找到第二大的数
	var ret []int

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			//注意: 这里先填满窗口的前 k - 1, 记录ret放在else分支中
			for len(rankQueue) > 0 && rankQueue[len(rankQueue)-1] < nums[i] {
				rankQueue = rankQueue[:len(rankQueue)-1]
			}
			rankQueue = append(rankQueue, nums[i])
		} else {
			// 窗口向前滑动，加入新数字
			for len(rankQueue) > 0 && rankQueue[len(rankQueue)-1] < nums[i] {
				rankQueue = rankQueue[:len(rankQueue)-1]
			}
			rankQueue = append(rankQueue, nums[i])
			// 记录当前窗口的最大值
			ret = append(ret, rankQueue[0])
			// 移出旧数字
			if nums[i-k+1] == rankQueue[0] {
				rankQueue = rankQueue[1:]
			}
		}
	}
	return ret
}
