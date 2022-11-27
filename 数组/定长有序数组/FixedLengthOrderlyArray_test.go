package 定长有序数组

import "testing"

func TestFixedLengthOrderlyArray_addItem(t *testing.T) {
	array := newFixedLengthOrderlyArray(4)
	array.addItem(1)
	array.addItem(10)
	array.addItem(5)
	array.addItem(3)
	array.addItem(2)
	array.addItem(10)
	array.addItem(5)
	array.addItem(3)
	array.addItem(2)
	array.addItem(10)
	array.addItem(5)
	array.addItem(3)
	array.addItem(22)
	array.addItem(10)
	array.addItem(5)
	array.addItem(3)
	array.addItem(21)
	array.addItem(32)
	array.addItem(12)
	array.addItem(10)
	array.addItem(5)
	array.addItem(36)
	array.addItem(32)
	array.addItem(10)
	t.Logf("%+v", array)
}
