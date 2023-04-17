/**
  @author: wangyingjie
  @since: 2023/4/17
  @desc:
**/

package main

type T struct{}

func (*T) foo() {
}

func (T) bar() {
}

type S struct {
	*T
}

func main() {
	s := S{}
	_ = s.foo
	s.foo()
	_ = s.bar //可以看成: (*s.T).bar("."优先级比*高), 因为要解引用T, *T是空指针, 所以报错
}
