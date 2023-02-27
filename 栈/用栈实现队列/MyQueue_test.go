/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc: //TODO
**/

package 用栈实现队列

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name    string
		operate []string
		data    [][]int
	}{
		{
			name:    "case",
			operate: []string{"Push", "Push", "Peek", "Pop", "Empty", "Pop"},
			data: [][]int{
				[]int{1}, []int{2}, []int{}, []int{}, []int{}, []int{},
			},
		},

		{
			name:    "case1",
			operate: []string{"Push", "Push", "Push", "Pop", "Pop", "Pop"},
			data: [][]int{
				[]int{1}, []int{2}, []int{3}, []int{}, []int{}, []int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myQueue := Constructor()
			refQueue := reflect.ValueOf(&myQueue)
			for i, funcName := range tt.operate {

				var values []reflect.Value
				for _, i3 := range tt.data[i] {
					values = append(values, reflect.ValueOf(i3))
				}

				ret := refQueue.MethodByName(funcName).Call(values)
				fmt.Println(funcName)
				for _, value := range ret {
					fmt.Println(value.Interface())
				}
			}
		})
	}
}
