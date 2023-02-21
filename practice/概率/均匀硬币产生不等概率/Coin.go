/**
  @author: wangyingjie
  @since: 2023/2/21
  @desc: https://leetcode.cn/problems/implement-rand10-using-rand7/solutions/427572/cong-pao-ying-bi-kai-shi-xun-xu-jian-jin-ba-zhe-da/
**/

package 均匀硬币产生不等概率

import (
	"math/rand"
	"time"
)

// coin
//  @Description: 硬币正反两面0.5
//  @return bool
func coin() int {
	rand.Seed(time.Now().UnixNano())
	if rand.Int()%2 == 1 {
		return 1
	} else {
		return 0
	}
}

// newCoin1divide4
//  @Description: 均匀概率的硬币 最终实现 P(0) = 1/4，P(1) = 3/4
//  @return int
func newCoin1divide4() int {
	//一个硬币分正面和反面, 0和1
	//抛两次的话则有 00, 01, 10, 11
	//抛3次则有000. 001, ......., 111
	//每次抛出相当于得到二进制的1位
	//抛出k次硬币后的数据范围为 [0, 2^k - 1]

	//依题意, 最终要求的概率为1/4和3/4, 那么至少需要4个数据, 选k=2
	//需要抛2次硬币, 数据范围[0, 3]是等概率的
	//那么随便挑选一种情况即为1/4
	first := coin()
	second := coin()
	if first == second && first == 0 {
		return 0
	} else {
		return 1
	}
}

// newCoin1divide3
//  @Description: 均匀概率的硬币 最终实现 P(0) = 1/3，P(1) = 2/3
//  @return int
func newCoin1divide3() int {
	//依题意, 最终要求的概率为1/3和2/3, 那么至少需要3个数据, 选k=2, 随便挑选一种情况即为1/4
	for true {
		//但是得满足1/3, 那么需要去掉一种情况, 这里去掉"11"这种, 当遇到这种情况时重新进行选择
		first := coin()
		second := coin()
		if first == second && first == 0 {
			return 0
		} else if first != second {
			return 1
		}
	}
	return -1
}

// newCoin3divide7
//  @Description: 均匀概率的硬币 最终实现 P(0) = 3/10，P(1) = 7/10
//  @return int
func newCoin3divide10() int {
	//依题意, 最终要求的概率为3/10和7/10, 那么至少需要10个数据, 选k=4作为抛次数, 范围[0, 15]中的数随机等概率1/16
	for true {
		//要满足总数为10, 那么需要去掉6种(16-10)情况
		//3/10: 就是[0, 2] => 0000, 0001, 0010
		//7/10: 就是[3, 9] => 0010, ......, 1001
		//其他情况: 重新进入循环
		num := 0
		for i := 0; i < 4; i++ {
			num = (num << 1) + coin()
		}
		if num <= 2 {
			return 0
		} else if num <= 9 {
			return 1
		}
	}
	return -1
}
