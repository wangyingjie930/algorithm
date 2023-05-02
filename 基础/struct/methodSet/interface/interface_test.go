/**
  @author: wangyingjie
  @since: 2023/4/30
  @desc:
**/

package interface_test

import (
	"algorithm/基础/struct/methodSet"
	"testing"
)

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

func TestInterface(test *testing.T) {
	var t T
	var pt *T
	methodSet.PrintMethodSet(&t)
	methodSet.PrintMethodSet(&pt)
}
