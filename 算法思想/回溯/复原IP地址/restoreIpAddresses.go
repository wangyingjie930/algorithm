/**
  @author: wangyingjie
  @since: 2023/3/22
  @desc: https://leetcode.cn/problems/restore-ip-addresses/description/
**/

package 复原IP地址

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}
	cur := ""
	var ret []string
	helper(cur, s, 0, &ret)
	return ret
}

// helper
//  @Description: 回溯3件套: 当前选择, 剩余选择, 结果保存
//  @param cur
//  @param remain
//  @param digitNum
//  @param ret
func helper(cur string, remain string, digitNum int, ret *[]string) {
	if digitNum == 4 && len(remain) == 0 {
		*ret = append(*ret, cur[1:])
		return
	} else if digitNum >= 4 {
		return
	}

	var digit string
	for i := 0; i < 3 && i < len(remain); i++ {
		//选取满足的ip整数, 加入ip字符串
		digit = digit + string(remain[i])
		num, _ := strconv.Atoi(digit)
		if digit == "0" || (num <= 255 && digit[0] != '0') {
			cur = cur + "." + digit

			helper(cur, remain[i+1:], digitNum+1, ret)

			cur = cur[0 : len(cur)-len(digit)-1]
		}
	}
}
