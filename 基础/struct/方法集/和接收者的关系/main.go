package main

import "fmt"

type Student struct {
	age  int8
	name string
}
type StudentPoint *Student

func (Student) sayHello() { //省略receiver 的参数参数名字
	fmt.Println("hello world")
}

func (s Student) showName() {
	fmt.Println(s.name)
}

func (s *Student) setName(newName string) {
	s.name = newName
}

/*func (s StudentPoint) showName2() { // Error：接受者(receiver)为指针类型
	fmt.Println(s.name)
}*/

// main
//  @Description: 实例value或pointer调用全部的方法，编译器会自动转换
func main() {
	s := Student{}
	s.setName("dq") //go会自动转为 (&s).setName("dq")

	var s2 = &s
	s2.showName() //o会自动转为 (*s2).showName()
}
