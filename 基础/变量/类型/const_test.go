/**
  @author: wangyingjie
  @since: 2023/4/17
  @desc:
**/

package alias_test

import (
	"fmt"
	"testing"
)

func TestConstDefine(t *testing.T) {
	const x = 123
	const y = 1.23 //常量定义后不使用也是可以的
	fmt.Println(x)
}

// TestConstDefineMulti
//  @Description: 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
//  @param t
func TestConstDefineMulti(t *testing.T) {
	const (
		x uint16 = 120
		y
		s = "abc"
		z
	)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
}
