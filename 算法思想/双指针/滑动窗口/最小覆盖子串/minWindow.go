/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc: https://leetcode.cn/problems/minimum-window-substring/description/
**/

package 最小覆盖子串

func minWindow(s string, t string) string {
	need := make(map[byte]int, len(t))
	window := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0 //满足条件的字符串长度(字符条件: 出现次数同t的出现次数一样)
	res := ""  //返回字符串
	for right < len(s) {
		if _, exist := need[s[right]]; exist {
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				valid++
			}
		}
		right++
		//满足条件的字符串长度和t的映射一样长时判断为覆盖
		for valid == len(need) {
			//左窗口收缩, 直到第不能覆盖
			if res == "" || len(res) > len(s[left:right]) {
				//在满足条件时获得最小子串
				res = s[left:right]
			}
			if _, exist := need[s[left]]; exist {
				if window[s[left]] == need[s[left]] {
					valid--
				}
				//减少出现次数
				window[s[left]]--
			}
			left++
		}
	}
	return res
}
