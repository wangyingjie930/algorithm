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

// Test_equal
//  @Description:
// 结构体只能比较是否相等，但是不能比较大小
//  @param t
func Test_equal(t *testing.T) {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		// 如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；
		fmt.Println("sn1 == sn2")
	}
}

// Test_sequence
//  @Description: 结构体属性的顺序不一致的话是不能进行比较
//  @param t
/*func Test_sequence(t *testing.T) {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		name string
		age  int
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
}

// Test_type
//  @Description:
//  struct 中有 bool、数值型、字符、指针、数组是可以进行比较的, 切片、map、函数等是不能比较的
//  @param t
/*func Test_type(t *testing.T) {
	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}*/

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
