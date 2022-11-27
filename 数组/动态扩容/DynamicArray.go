package 动态扩容

type DynamicArray struct {
	array []interface{} //定长数组
	size int //当前保存的个数
}

//初始化固定长度的数组
func newDynamicArray(size int) *DynamicArray {
	return &DynamicArray{array: make([]interface{}, size), size: 0}
}

//扩容
func (array *DynamicArray) resize() {
	newArray := make([]interface{}, array.size * 2)
	for i, val := range array.array {
		newArray[i] = val
	}
	array.array = newArray
}

//当固定长度等于当前保存个数则进行扩容
func (array *DynamicArray) addItem(item interface{}) {
	if len(array.array) == array.size {
		array.resize()
	}
	array.array[array.size] = item
	array.size ++
}