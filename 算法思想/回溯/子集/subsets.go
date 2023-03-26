/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc: https://leetcode.cn/problems/subsets/description/
**/

package 子集

func subsets(nums []int) [][]int {
	var ret [][]int
	var cur []int
	helper(cur, nums, &ret)
	return ret
}

func helper(cur []int, remain []int, ret *[][]int) {
	//注意这里不需要终止条件
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	*ret = append(*ret, tmp)

	for i := 0; i < len(remain); i++ {
		cur = append(cur, remain[i])
		helper(cur, remain[i+1:], ret)
		cur = cur[:len(cur)-1]
	}
}
