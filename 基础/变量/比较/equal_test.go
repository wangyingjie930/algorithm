/**
  @author: wangyingjie
  @since: 2023/4/17
  @desc:
**/

package equal_test

import (
	"fmt"
	"testing"
)

// Test_equal
//  @Description:
// 结构体只能比较是否相等，但是不能比较大小
//  @param t
func TestStruct(t *testing.T) {
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

func TestArray(t *testing.T) {
	a := [2]int{5, 6}
	b := [2]int{5, 7}
	////如果b := [3]int{5, 7, 8}进行比较会发生报错,数组的长度也是数组类型的组成部分，所以 a 和 b 是不同的类型，是不能比较的
	if a == b {
		//数组也是只能比较是否相等, 并不能比较>或<
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

func TestString(t *testing.T) {
	//str == nil //error 字符串不能和nil进行比较
}

func TestSlice(t *testing.T) {
	//fmt.Println([...]int{1} == [2]int{1}) //error, 不同类型不能进行比较, 包括数组长度不一样也是不能比较的
	//fmt.Println([]int{1} == []int{1}) //error, 切片不能进行比较, 切片、map、函数等是不能比较的
	//切片、map、函数等是不能比较的, 只能判断是否==nil
	a := []int{}
	b := map[int]int{}
	fmt.Println(a == nil, b == nil)
}

func TestInterface(t *testing.T) {
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = x == y //相当于切片和nil进行比较, ok
	_ = y == y //相当于两个切片进行比较, error
}
