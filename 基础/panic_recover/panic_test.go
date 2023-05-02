/**
  @author: wangyingjie
  @since: 2023/5/2
  @desc:
**/

package panic_recover

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		defer func() {
			recover()
		}()
	}()
	panic(1)
}

func TestPanic1(t *testing.T) {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer fmt.Print(recover())
		panic(1)
	}()
	panic(2)
}
