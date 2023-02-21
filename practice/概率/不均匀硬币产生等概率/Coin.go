/**
  @author: wangyingjie
  @since: 2023/2/21
  @desc: 一枚不均匀的硬币 coin()，能够返回 0、1 两个值，其概率分别为 0.6、0.4。要求使用这枚硬币，产生均匀的概率分布。
  @see: https://leetcode.cn/problems/implement-rand10-using-rand7/solutions/427572/cong-pao-ying-bi-kai-shi-xun-xu-jian-jin-ba-zhe-da/
**/

package 不均匀硬币产生等概率

import (
	"math/rand"
	"time"
)

// coin
//  @Description: 硬币以0.6的概率向上(1), 0.4的概率向下(0)
//  @return bool
func coin() int {
	rand.Seed(time.Now().UnixNano())
	if rand.Int()%10 > 5 {
		return 1
	} else {
		return 0
	}
}

// averageCoin
//  @Description:
// 00: p(0) * p(0)
// 01: p(0) * p(1)
// 10: p(1) * p(0)
// 11: p(1) * p(1)
// 因为01和10的概率一样, 所以以这两种情况作为概率的依据返回
// 遇到00或11则重新执行
//  @return bool
func averageCoin() int {
	for true {
		first := coin()
		if coin() != first {
			//遇到01或10则返回
			return first
		}
	}
	return -1
}
