/**
  @author: wangyingjie
  @since: 2023/4/17
  @desc:
**/

package count_test

import (
	"fmt"
	"testing"
)

type T struct {
	x int
	y *int
}

func TestT(t1 *testing.T) {
	i := 20
	t := T{x: 10, y: &i}

	p := &t.x //.优先级大于&,*, 得到&10

	*p++ //&,*优先级大于++, -- //11
	*p-- //10

	t.y = p

	fmt.Println(*t.y) //输出: 10
}
