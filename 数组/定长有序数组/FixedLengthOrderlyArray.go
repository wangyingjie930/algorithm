package 定长有序数组

type FixedLengthOrderlyArray struct {
	array []int //定长数组
	size int //当前保存的个数
}

//初始化固定长度的数组
func newFixedLengthOrderlyArray(size int) *FixedLengthOrderlyArray {
	return &FixedLengthOrderlyArray{array: make([]int, size), size: 0}
}

func (array *FixedLengthOrderlyArray) addItem(item int) {
	if len(array.array) == array.size {
		return
	}
	//在有序数组中通过二分法找到适合插入的索引
	index := array.binarySearch(0, array.size - 1, item)
	//迁移位置, 给index腾出位置
	for i := array.size; i > index; i-- {
		array.array[i] = array.array[i - 1]
	}
	//插入
	array.array[index] = item
	array.size ++
}

//二分查找
func (array *FixedLengthOrderlyArray) binarySearch(start, end int, item int) int {

	if array.size == 0 {
		return 0
	}

	if start == end {
		if array.array[start] >= item {
			return start
		}else {
			return start + 1
		}
	}

	mid := (start + end) / 2
	if array.array[mid] < item {
		return array.binarySearch(mid + 1, end, item)
	}else {
		return array.binarySearch(start, mid, item)
	}
}