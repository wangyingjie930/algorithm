/**
  @author: wangyingjie
  @since: 2023/4/30
  @desc:
**/

package alias

import (
	"algorithm/基础/struct/methodSet"
	"testing"
)

type T struct{}

func (T) M1()  {}
func (*T) M2() {}

type Interface interface {
	M1()
	M2()
}

type NewT T
type NewInterface Interface

type AliasT = T
type AliasInterface = Interface

func TestAlias(test *testing.T) {
	var t T                //M1
	var pt *T              //M1 M2
	var it Interface       //M1 M2
	var newt NewT          //empty
	var newpt *NewT        //empty
	var newit NewInterface //M1 M2

	var aliast AliasT          //M1
	var aliaspt *AliasT        //M1 M2
	var aliasit AliasInterface //M1 M2

	methodSet.PrintMethodSet(&t)
	methodSet.PrintMethodSet(&pt)
	methodSet.PrintMethodSet(&it)

	methodSet.PrintMethodSet(&newit)
	methodSet.PrintMethodSet(&newt)
	methodSet.PrintMethodSet(&newpt)

	methodSet.PrintMethodSet(&aliast)
	methodSet.PrintMethodSet(&aliaspt)
	methodSet.PrintMethodSet(&aliasit)
}
