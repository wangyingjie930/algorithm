package 跳表

import (
	"math"
	"math/rand"
)

const (
	// MAX_LEVEL 最高层数
	MAX_LEVEL = 16
)

//跳表节点结构体
type skipListNode struct {
	//跳表保存的值
	v interface{}
	//用于排序的分值
	score int
	//层高
	level int
	//每层前进指针
	forwards []*skipListNode
}

func newSkipListNode(v interface{}, score int, level int) *skipListNode {
	return &skipListNode{v: v, score: score, level: level, forwards: make([]*skipListNode, level, level)}
}

// SkipList 跳表结构体
type SkipList struct {
	//跳表头结点
	head *skipListNode
	//跳表当前层数
	level int
	//跳表长度
	length int
}

func NewSkipList() *SkipList {
	//头结点，便于操作
	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL)
	return &SkipList{head, 1, 0}
}

// Insert 插入节点到跳表中
func (sl *SkipList) Insert(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	//查找插入位置
	cur := sl.head
	//记录新插入的节点的第i层应该连接哪个节点, 成为它的后继节点
	//key为层数, value为指向那个节点
	update := [MAX_LEVEL]*skipListNode{}
	i := MAX_LEVEL - 1
	for ; i >= 0; i-- {
		//从头节点的最顶层开始逐层遍历
		for nil != cur.forwards[i] {
			//在第i层时, 有可以跳跃的目标则进行判断, 成功则进行跳跃
			//跳跃的对象是整个节点, 并不是某一层
			if cur.forwards[i].v == v {
				//如果找到则返回
				return 2
			}
			if cur.forwards[i].score > score {
				//当前节点的跳跃节点的score>要查找的score, 记录当前的节点
				//并且留在当前节点不进行跳跃, 而是在当前节点往下一层继续查看下一个跳跃节点是否满足<score
				update[i] = cur
				break
			}
			//如果满足<score, 则进行跳跃
			//注意这是跳往整个节点的, 并不是跳往节点某一层
			cur = cur.forwards[i]
		}
		//当没有可以跳跃访问的节点, 默认取当前遍历的节点作为后续插入节点的头结点
		//也就是说update数组的每一项都是有数据的
		if nil == cur.forwards[i] {
			update[i] = cur
		}
	}

	//通过随机算法获取该节点层数
	//索引的动态更新
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	//创建一个新的跳表节点
	newNode := newSkipListNode(v, score, level)

	//原有节点连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	//如果当前节点的层数大于之前跳表的层数
	//更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	//更新跳表长度
	sl.length++

	return 0
}

// Find 查找
func (sl *SkipList) Find(v interface{}, score int) *skipListNode {
	if nil == v || sl.length == 0 {
		return nil
	}

	cur := sl.head
	//从当前跳表的最大层数开始遍历
	for i := sl.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			//进行跳跃
			cur = cur.forwards[i]
		}
	}

	return nil
}

