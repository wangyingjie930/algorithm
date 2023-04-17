/**
  @author: wangyingjie
  @since: 2023/4/17
  @desc: 当指针值赋值给变量或者作为函数参数传递时，会立即计算并复制该方法执行所需的接收者对象，与其绑定，以便在稍后执行时，能隐式第传入接收者参数
**/

package method_test

import (
	"fmt"
	"testing"
)

type N int

func (n N) test() {
	fmt.Println(n)
}

func TestMethod(t *testing.T) {
	var n N = 10
	p := &n

	n++
	f1 := n.test //此时复制的是n的值, n是11

	n++
	f2 := p.test //此时复制的是p的值. p是12

	n++
	fmt.Println(n)

	f1()
	f2()
	//输出: 13 11 12
}

type N1 int

func (n *N1) test() {
	fmt.Println(*n)
}

func TestMethod1(t *testing.T) {
	type N int

	var n N1 = 10
	p := &n

	n++
	f1 := n.test //此时n复制的是地址

	n++
	f2 := p.test //此时p复制的是地址

	n++
	fmt.Println(n)

	f1() //此时的地址存放的是13这个值
	f2()
	//输出: 13 13 13
}
