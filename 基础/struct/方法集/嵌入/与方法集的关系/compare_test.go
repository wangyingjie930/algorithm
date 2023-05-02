/**
  @author: wangyingjie
  @since: 2023/4/29
  @desc:
**/

package method_set

import (
	"fmt"
	"testing"
)

type T struct {
	a int
}

func (receiver T) func1() {
	receiver.a = 1
	fmt.Println("T::func1")
}

func (receiver *T) func2() {
	receiver.a = 2
	fmt.Println("*T::func2")
}

type S1 struct {
	T
}

type S2 struct {
	*T
}

// TestA
//  @Description: 也就是说编译器无论怎么样都能帮你自动安排好
// 当进行调用的时候是不需要进行方法集的判断
//  @param t
func TestS1Normal(t *testing.T) {
	s1 := S1{}
	s1.func1()
	fmt.Println(s1.a)
	s1.func2()
	fmt.Println(s1.a)
}

func TestS1Refer(t *testing.T) {
	s1 := &S1{}
	s1.func1()
	fmt.Println(s1.a)
	s1.func2()
	fmt.Println(s1.a)
}

func TestS2Normal(t *testing.T) {
	s2 := S2{T: new(T)}
	s2.func1()
	fmt.Println(s2.a)
	s2.func2()
	fmt.Println(s2.a)
}

func TestS2Refer(t *testing.T) {
	s2 := &S2{T: new(T)}
	s2.func1()
	fmt.Println(s2.a)
	s2.func2()
	fmt.Println(s2.a)
}
