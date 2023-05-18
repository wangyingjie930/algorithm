/**
  @author: wangyingjie
  @since: 2023/5/18
  @desc: https://leetcode.cn/problems/k-th-smallest-in-lexicographical-order/description/
**/

package findKthNumber

//findKthNumber
//@Description: 多路查找树/字典序的第K小数字/PNG图像.png
//@param n int
//@param k int
//@return int
func findKthNumber(n int, k int) int {
	cur := 1
	k--

	for k > 0 {
		steps := getSteps(cur, n)
		if steps <= k {
			k = k - steps
			cur = cur + 1
		} else {
			k--
			cur = cur * 10
		}
	}
	return cur
}

func getSteps(cur, n int) int {
	left, right := cur, cur
	steps := 0
	for left <= n {
		steps = min(right, n) - left + steps + 1
		left = left * 10
		right = right*10 + 9
	}
	return steps
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
