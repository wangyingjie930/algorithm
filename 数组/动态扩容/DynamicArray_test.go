package 动态扩容

import "testing"

func TestDynamicArray_addItem(t *testing.T) {
	array := newDynamicArray(2)
	array.addItem(1)
	array.addItem(2)
	array.addItem(2)
	array.addItem(2)
	array.addItem(2)
	array.addItem(2)
	t.Logf("%+v", array)
	t.Log(len(array.array))
}
