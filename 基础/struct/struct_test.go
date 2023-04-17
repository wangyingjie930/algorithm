/**
  @author: wangyingjie
  @since: 2023/4/15
  @desc: https://www.topgoer.cn/docs/gomianshiti/mian5
**/

package struct_test

import (
	"fmt"
	"testing"
)

type Integer int

//值类型接受者: 修改操作只是针对副本，无法修改接收者变量本身
func (a Integer) Add(b Integer) Integer {
	return a + b
}

/*
@see: 基础/miscellaneous.md:79 方法集的概念

//指针类型接收者:调用方法时可以修改接收者指针的任意成员变量
func (a *Integer) Add(b Integer) Integer {
	return *a + b
}*/

func TestAssert(t *testing.T) {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	ret := i.(*Integer)
	sum := ret.Add(b) //使用值接收者或者指针接收者都能调用
	fmt.Println(sum)
}

type X struct{}

func (x *X) test() {
	println(x, x == nil) //nil, true
}
func TestX(t *testing.T) {
	var a *X
	a.test() //相当于test(a)
	(&X{}).test()
	//X{}.test() //error, 不可直接寻址
}
