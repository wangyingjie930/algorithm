/**
  @author: wangyingjie
  @since: 2023/4/16
  @desc:
**/

package main

import (
	"fmt"
	"testing"
)

type Person struct {
	age  int8
	name string
}

func (s Person) showName() {
	fmt.Println(s.name)
}

func (s *Person) setName(newName string) {
	s.name = newName
}

// Student1
// @Description: 值嵌入
type Student1 struct {
	Person //Student1包含了Person,那么Student1对应的value和pointer包含Person
}

// Student2
// @Description: 指针嵌入
type Student2 struct {
	*Person
}

func TestInit(t *testing.T) {
	// 内嵌类型 Person默认值为 Person{age:0, name:""}
	s1 := Student1{}
	s1.setName("student1_01") // ok
	s1.showName()

	// 内嵌类型 *Person默认值为 nil
	s2 := &Student2{}
	s2.setName("student1_02") //Error，由于目前内嵌类型的值为nil，会触发报错
	s2.showName()

	// 给嵌入类型一个复制，就ok了
	s3 := &Student2{Person: &Person{age: 3, name: "s3"}}
	//s3 := &Student2{&Person{age:3, name:"s3"}} 和上一行等价
	s3.showName()
}

type Param map[string]interface{}

type Show struct {
	*Param
}

func TestInit1(t *testing.T) {
	s := new(Show)
	p := make(Param) //需要先初始化, 指针嵌入默认是nil的
	p["day"] = 2
	s.Param = &p
	fmt.Println((*s.Param)["day"])
}
