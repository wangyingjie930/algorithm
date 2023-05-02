/**
  @author: wangyingjie
  @since: 2023/4/30
  @desc:
**/

package methodSet

import (
	"fmt"
	"reflect"
)

func PrintMethodSet(i interface{}) {
	elemTyp := reflect.TypeOf(i)
	elemTyp = elemTyp.Elem()

	n := elemTyp.NumMethod()
	if n == 0 {
		fmt.Println(elemTyp, "'s method set is empty!")
		return
	}

	fmt.Println(elemTyp, "'s method set:")
	for j := 0; j < n; j++ {
		fmt.Println("--", elemTyp.Method(j).Name)
	}
	fmt.Println()
}
