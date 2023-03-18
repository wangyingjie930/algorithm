/**
  @author: wangyingjie
  @since: 2023/3/18
  @desc: https://leetcode.cn/problems/add-strings/description/
  掌握程度: todo: 具体编码很低, 重看
**/

package 字符串相加

import "strconv"

// addStrings
//  @Description: 从后往前遍历, 模拟进位相加运算
//  @param num1
//  @param num2
//  @return string
func addStrings(num1 string, num2 string) string {
	result := ""
	add := 0
	for i, j := len(num1)-1, len(num2)-1;
	//位数或者还有进位时, 都要执行相加操作
	i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {

		a3 := 0
		if i >= 0 {
			a3 = a3 + int(num1[i]-'0') //关于字符数字转成数字的方法
		}
		if j >= 0 {
			a3 = a3 + int(num2[j]-'0')
		}
		a3 = a3 + add

		result = strconv.Itoa(a3%10) + result //前后拼接成数字字符串
		add = a3 / 10                         //获得下个位数相加的进位
	}
	return result
}
