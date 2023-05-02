/**
  @author: wangyingjie
  @since: 2023/4/30
  @desc:
**/

package enmbed

import (
	"algorithm/基础/struct/methodSet"
	"testing"
)

type T1 struct{}

func (T1) T1M1()   {}
func (*T1) PT1M2() {}

type T2 struct{}

func (T2) T2M1()   {}
func (*T2) PT2M2() {}

type T struct {
	T1
	*T2
}

func TestEmbed(test *testing.T) {
	var t T
	var pt *T
	methodSet.PrintMethodSet(&t)  //PT2M2 T1M1 T2M1
	methodSet.PrintMethodSet(&pt) //PT1M2 PT2M2 T1M1 T2M1
}
