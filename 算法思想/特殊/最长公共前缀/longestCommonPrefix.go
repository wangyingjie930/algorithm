/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc: https://leetcode.cn/problems/longest-common-prefix/description/?orderBy=most_votes
**/

package 最长公共前缀

// longestCommonPrefix
//  @Description: 没啥原理, 就是垂直遍历
//  @param strs
//  @return string
func longestCommonPrefix(strs []string) string {
	ret := ""

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return ret
			}
		}
		ret = ret + string(strs[0][i])
	}
	return ret
}
