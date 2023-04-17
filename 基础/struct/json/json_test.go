/**
  @author: wangyingjie
  @since: 2023/4/17
  @desc:
**/

package json_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	js := `{"name":"user"}`
	var p struct {
		//因为 name 首字母是小写，导致其他包不能访问，所以输出为空结构体
		name string `json:"name"`
		//需要改为: Name string `json:"name"`
	}

	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(p.name)
}
