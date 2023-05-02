/**
  @author: wangyingjie
  @since: 2023/4/29
  @desc: @Description: http://www.findme.wang/blog/detail/id/559.html

**/

package method_set

import "fmt"

type Human interface { //定义接口
	showName()
	setName(string)
}

type Person struct {
	age  int8
	name string
}

func (s Person) showName() { // Person和*Person都实现了Human的showName方法
	fmt.Println(s.name)
}
func (s *Person) setName(newName string) { // 只有*Person实现了Human的setNanme方法
	s.name = newName
}

//所以*Person实现了Human接口, Person没有实现Human接口

type Student1 struct { // 以值类型，嵌入
	Person
}

type Student2 struct { // 以指针类型，嵌入
	*Person
}

// main
// 前提: 在进行接口断言的时候才需要考虑方法集, 平常的调用由编译器自动转换所以不用考虑方法集
//  @Description: 规则如下:
//1、类型 S 包含匿名字段 T，则 S和*S 方法集包含 T 方法。
//2、类型 S 包含匿名字段 *T，则 S和 *S 方法集包含 T + *T 方法。
//3、不管嵌入的是T还是*T，*S方法集，包含 T + *T 方法。
func testInterface() {
	//Student1只能实现Person的方法[规则1]
	//*Student1可以实现Person和*Person的方法 [规则3]
	//*Student2/Student2可以实现Person和*Person的方法[规则2]

	//总结: 只要是*S都能实现T+*T的方法
	//S{T}时, S只能实现T的方法

	/**var s1 Human = Student1{} //报错:Student1 does not implement Human (setName method has pointer receiver)
	s1.setName("student1_01")
	s1.showName()*/

	var s2 Human = &Student1{} //ok 应用第1条和弟3条规则
	s2.setName("student1_02")
	s2.showName()

	var s3 Human = Student2{&Person{}} //ok ，应用第2条规则
	s3.setName("student2_01")
	s3.showName()

	var s4 Human = &Student2{&Person{}} //ok ，应用第2条规则
	s4.setName("student2_02")
	s4.showName()
}
