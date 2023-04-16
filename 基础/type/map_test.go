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

func Test_map(t *testing.T) {
	var mapper map[int]int
	//打印一个 map 中不存在的值时，返回元素类型的零值
	fmt.Print(mapper[1])

	delete(mapper, 1) //删除 map 不存在的键值对时，不会报错
}

func TestRangeMap(t *testing.T) {
	var m = map[string]int{
		"tony": 21,
		"tom":  22,
		"jim":  23,
	}
	for _ = range m {
		//指针副本也指向真实map，因此for range操作均操作的是源map
		delete(m, "tony")
	}
	t.Log(m)
}
