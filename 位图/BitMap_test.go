package 位图

import (
	"fmt"
	"testing"
)

func TestNewBitMap(t *testing.T) {
	bitmap := NewBitMap(64)
	bitmap.set(45)
	fmt.Println(bitmap.get(45), bitmap.get(42))
	bitmap.remove(45)
	fmt.Println(bitmap.get(45), bitmap.get(42))
}
