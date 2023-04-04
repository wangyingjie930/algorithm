/**
  @author: wangyingjie
  @since: 2023/4/4
  @desc:
**/

package 面试汇总

/**
一个有序数组，找出连续x个元素之和等于y的子数组，比如有序数组为[1，2，2，3，4，5，6，7，8]，连续元素求和等于7的子数组是[2,2,3],[3,4],[7]
*/

func test(nums []int) [][]int {
	left, right := 0, 0
	remain := 7

	var window []int
	var ret [][]int

	for right < len(nums) {
		c := nums[right]
		right++
		window = append(window, c)
		remain = remain - c

		for remain < 0 {
			window = window[1:]
			remain = remain + nums[left]
			left++
		}

		if remain == 0 {
			ret = append(ret, window)
		}
	}

	return ret
}
