/**
  @author: wangyingjie
  @since: 2023/4/16
  @desc: 类型别名
**/

package alias_test

import (
	"fmt"
	"testing"
)

type MyInt1 int   //创建新的类型
type MyInt2 = int //类型的别名

func TestAlias(t *testing.T) {
	var i int = 0
	//var i1 MyInt1 = i //int和MyInt1类型不匹配, 不能直接赋值, 需要强转
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i //因为是别名, 所以类型匹配
	fmt.Println(i1, i2)
}

const i = 100

var j = 123

// TestConst
//  @Description: 常量无法寻址
//  @param t
func TestConst(t *testing.T) {
	fmt.Println(&j, j)
	//fmt.Println(&i, i) //error, 常量无法寻址
	//常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用
}

type User struct{}
type User1 User
type User2 = User //这里是别名

func (i User1) m1() {
	fmt.Println("m1")
}
func (i User) m2() {
	fmt.Println("m2")
}

func TestUser(t *testing.T) {
	var i1 User1
	var i2 User2
	i1.m1()
	i2.m2() //因为User2是User的别名, 所以User实现了, User2也就实现了
}

// TestByteAlias
//  @Description: byte的别名是uint8,rune的别名是int32
//  @param t
func TestByteAlias(t *testing.T) {
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b
	func(x byte) {
		fmt.Println(x)
	}(c)
}
