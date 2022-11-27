package consistentTest

import (
	"hash/crc32"
	"sort"
)

//声明新切片类型
type units []uint32

//返回切片长度
func (x units) Len() int {
	return len(x)
}

//比对两个数大小
func (x units) Less(i, j int) bool {
	return x[i] < x[j]
}

//切片中两个值的交换
func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type Consistent struct {
	circle map[uint32]string
	sortedHashes units
}

func NewConsistent() *Consistent {
	return &Consistent{
		//初始化变量
		circle: make(map[uint32]string),
	}
}

func (c *Consistent) hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *Consistent) search(key uint32) int {
	//使用二分查找
	index := sort.Search(len(c.sortedHashes), func(i int) bool {
		return c.sortedHashes[i] >= key
	})
	//超过数组长度返回第一个索引
	if index > len(c.sortedHashes) {
		index = 0
	}
	return index
}

func (c *Consistent) Get(name string) (string, error) {
	//先哈希
	key := c.hashKey(name)
	//再查找
	index := c.search(key)
	//排序键->排序值->哈希环值
	return c.circle[c.sortedHashes[index]], nil
}

//更新排序，方便查找
func (c *Consistent) updateSortedHashes() {
	var hashes units
	//添加hashes
	for k := range c.circle {
		hashes = append(hashes, k)
	}

	//对所有节点hash值进行排序，
	//方便之后进行二分查找
	sort.Sort(hashes)
	//重新赋值
	c.sortedHashes = hashes
}

//添加节点
func (c *Consistent) Add(element string) {
	//循环虚拟节点，设置副本
	c.circle[c.hashKey(element)] = element
	//更新排序
	c.updateSortedHashes()
}