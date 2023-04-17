/**
  @author: wangyingjie
  @since: 2023/4/16
  @desc:
**/

package type_test

import (
	"fmt"
	"testing"
)

func Test_String(t *testing.T) {
	str := "hello"
	//str[0] = 'x' //Go 语言中的字符串是只读的, 会报错的
	fmt.Println(str)
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return []string{"North", "East", "South", "West"}[d]
}

func TestConvertString(t *testing.T) {
	//如果类型定义了 String() 方法，当使用 fmt.Printf()、fmt.Print() 和 fmt.Println() 会自动使用 String() 方法，实现字符串的打印
	fmt.Println(South) //输出: "South"
}
