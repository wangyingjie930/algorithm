/**
  @author: wangyingjie
  @since: 2023/2/24
  @desc: https://leetcode.cn/problems/permutations/
**/

package 全排列

func permute(nums []int) [][]int {
	used := make(map[int]bool, len(nums))
	var cur []int
	var ret [][]int
	dfs(nums, used, cur, &ret)
	return ret
}

// dfs
//  @Description:
//  @param nums
//  @param used 是否已经使用完【选择列表】
//  @param cur 【路径】
//  @param ret 引用传递改变值
func dfs(nums []int, used map[int]bool, cur []int, ret *[][]int) {
	if len(cur) == len(nums) {
		//【结束条件：选择列表为空】
		tmp := make([]int, len(nums))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			//标记使用
			used[i] = true
			//标记走过的路径
			cur = append(cur, nums[i])
			dfs(nums, used, cur, ret)
			//还原标记
			used[i] = false
			//还原路径
			cur = cur[0 : len(cur)-1]
		}
	}
	return
}
