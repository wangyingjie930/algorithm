/**
  @author: wangyingjie
  @since: 2023/4/15
  @desc: http://c.biancheng.net/view/4118.html
**/

package forRange_test

import (
	"fmt"
	"testing"
)

func Test_error(t *testing.T) {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		//这里的val是重新分配出来的变量, 并且循环过程中这个val的地址一直不变, 只有里面存放的值在不断地变
		//所以这里的m[key] = &val得到的是val的地址, 并没有指向slice
		//所以结果是所有的元素都拿到了val的地址, 每个元素拿到同一份数据, 遍历结束后val存放的值是切片最后的元素
		m[key] = &val
		fmt.Printf("%+v\n", &val)
	}

	for k, v := range m {
		/**
		输出:
			0 -> 3
			1 -> 3
			2 -> 3
			3 -> 3
		*/
		fmt.Println(k, "->", *v)
	}
}

// Test_true
//  @Description: 正确写法
//  @param t
func Test_true(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		value := val //表示每遍历一次重新分配一个新的value
		fmt.Printf("%v\n", &value)

		//每个元素可以指向不同地址的value
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "===>", *v)
	}
}

// TestRangeNil
//  @Description:
// 对nil map和slice的循环次数将是0
// 对nil数组的循环次数将取决于它的数组类型定义的长度
// 对nil channel的range操作将永远阻塞当前goroutine
//  @param t
func TestRangeNil(t *testing.T) {
	for range []int(nil) { //循环次数将是0
		fmt.Println("Hello")
	}

	for range map[string]string(nil) { //循环次数将是0
		fmt.Println("world")
	}

	for i := range (*[5]int)(nil) {
		fmt.Println(i) // 0 1 2 3 4
	}

	for range chan bool(nil) { // block here
		fmt.Println("Bye") //fatal error: all goroutines are asleep - deadlock!
	}
}
