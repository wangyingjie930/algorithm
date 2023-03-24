/**
  @author: wangyingjie
  @since: 2023/3/24
  @desc: https://leetcode.cn/problems/multiply-strings/description/

**/

package 字符串乘法计算

import "strconv"

func multiply(num1 string, num2 string) string {
	if num2 == "0" || num1 == "0" {
		return "0"
	}

	res := make([]int, len(num2)+len(num1))
	for i := len(num2) - 1; i >= 0; i-- {
		for j := len(num1) - 1; j >= 0; j-- {
			mul := int(num2[i]-'0') * int(num1[j]-'0')

			//i可以看做是起始位置, j可以看成从i开始走几步
			//从规律可以知道最后一位是i+j+1
			sum := mul + res[i+j+1]
			res[i+j+1] = sum % 10
			res[i+j] = res[i+j] + sum/10
		}
	}

	i := 0
	for i < len(res) && res[i] == 0 {
		i++
	}

	ret := ""
	for ; i < len(res); i++ {
		ret = ret + strconv.Itoa(res[i])
	}
	return ret
}
