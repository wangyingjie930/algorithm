/**
  @author: wangyingjie
  @since: 2023/5/2
  @desc:
**/

package interface_test

import (
	"fmt"
	"testing"
)

type I interface {
	Value() int
}

type A int

func (a A) Value() int {
	return int(a)
}

func TestInterface(t *testing.T) {
	var i I = A(1)
	var a A = 1
	fmt.Println(
		i.Value(), //实际上等价于A.Value(i), 但是因为这个i是接口类型而没有具体的值, 所以栈上是无法分配空间的
		//于是, 改成传递指针的方式, (*A).Value(&i)即可,
		//所以A类型实现了Value的同时, 编译器帮忙实现了*A的Value函数
		A.Value(a),
		(*A).Value(&a),
	)
}
