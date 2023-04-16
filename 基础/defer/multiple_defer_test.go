/**
  @author: wangyingjie
  @since: 2023/4/15
  @desc:
**/

package defer_test

import (
	"fmt"
	"testing"
)

func TestMultiple(t *testing.T) {
	a := 1
	b := 2
	//1. 执行calc(“10”,1,2)
	//4. 执行calc("1", 1, 3)
	defer calc("1", a, calc("10", a, b))
	a = 0
	//2. 执行calc(“20”,0,2)
	//3. 执行calc(“2”,0,2)
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
