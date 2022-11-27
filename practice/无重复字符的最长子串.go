package practice

/**
无重复字符的最长子串

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。


关键字：重复字符 -> 出现1次
-模式识别1：一旦涉及出现次数，需要用到散列表-构造子串，散列表存下标
-模式识别2：涉及子串，考虑滑动窗口
 */

func LengthOfLongestSubstring(s string) int {
	/**
	map表示遇到当前字符start指针应该从哪里开始才能确保start到当前字符之间是没有重复的,也就是窗口收缩的程度
	 */
	map1 := make(map[string]int)
	start, end := 0, 0
	ans := 0
	for start < len(s) && end < len(s) {
		//表示如果发现有重复的字符时, 进行窗口的收缩
		//收缩是依照排除之前的重复字符为依据(map对应存放的位置)
		if v, found := map1[string(s[end])]; found && v > start {
			start = v
		}
		if ans < end - start + 1 {
			ans = end - start + 1
		}
		map1[string(s[end])] = end + 1
		end ++
	}
	return ans
}
