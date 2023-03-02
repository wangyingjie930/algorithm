/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc: https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
**/

package 无重复字符的最长子串

// lengthOfLongestSubstring
//  @Description: 很显然求子串可以尝试滑动窗口, 模板走起
//  @param s
//  @return int
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0          //模板: 左右指针
	window := make(map[byte]int) // 模板: 维护一个窗口

	maxLen := 0
	hasRepeat := 0
	for right < len(s) {
		//明确好扩展/收缩时需要维护哪些变量
		if v, exist := window[s[right]]; exist && v >= 1 {
			hasRepeat++
		}
		window[s[right]]++
		right++

		for hasRepeat > 0 {
			//明确好什么时候应该收缩: 就是窗口中包含重复字符
			window[s[left]]--
			if window[s[left]] == 1 {
				//当有字符缩减为一个时表示这个字符已经不需要记录到重复统计里面
				hasRepeat--
			}
			left++
		}
		if maxLen < len(s[left:right]) {
			maxLen = len(s[left:right])
		}
	}
	return maxLen
}
