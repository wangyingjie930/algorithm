/**
  @author: wangyingjie
  @since: 2023/2/14
  @desc: 将数组打乱
**/

package 洗牌算法

import (
	"math/rand"
	"time"
)

func shuffle(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		//伪随机数: 就是通过一套固定的算法使人看起来随机, 但其实是有规律可循只是比较复杂而已, 并不是物理意义上的真随机
		//随机种子相当于算法的初始值, 之后通过随机算法生成新的值
		//相同的随机种子(初始值)一定会产生相同的新值
		//所以每次调用需要设置新的种子, 不然每次的随机数都是一样的
		rand.Seed(time.Now().Unix())
		random := rand.Int()
		index := random % i
		nums[index], nums[i] = nums[i], nums[index]
	}
}
