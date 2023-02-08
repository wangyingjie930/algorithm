/**
  @author: wangyingjie
  @since: 2023/1/14
  @desc: btree的添加操作
**/

package BTree

import "sort"

// B树节点
type node struct {
	items    items // 存放的元素数组
	children children //孩子节点数组
}

// BTree B树
type BTree struct {
	degree int //除 root 和叶子节点外，其它节点至少有 degree+1 个子节点，最多有 2degree+1 个子节点
	length int
	root   *node
}

// b树中节点的元素
type items []Item

// Item 节点中的元素必须实现一个相互比较的接口
type Item interface {
	Less(than Item) bool
}

// b树中的孩子节点
type children []*node

// ReplaceOrInsert B树插入操作
func (t *BTree) ReplaceOrInsert(item Item) Item {
	if item == nil {
		panic("nil item being added to BTree")
	}
	if t.root == nil {
		//当树为空节点的时候
		t.root = new(node)
		t.root.items = append(t.root.items, item)
		t.length++
		return nil
	} else {
		if len(t.root.items) >= 2 * t.degree + 1 {
			//如果根节点已满, 则进行分裂, 返回的是新的节点和中间待上升的元素
			item2, second := t.root.split((2 * t.degree + 1) / 2)
			oldroot := t.root
			//新的节点作为根节点
			t.root = new(node)
			//将待上升的元素加入
			t.root.items = append(t.root.items, item2)
			//老的节点作为孩子
			t.root.children = append(t.root.children, oldroot, second)
		}
	}
	//从根节点开始进行插入元素操作
	out := t.root.insert(item, 2 * t.degree + 1)
	if out == nil {
		t.length++
	}
	return out
}

//节点插入操作
func (n *node) insert(item Item, maxItems int) Item {
	//当前节点是否含有该元素, 使用二分查找法
	//若找到, 返回找到的索引, 若找不到, 返回的是中间位置的索引
	i, found := n.items.find(item)
	if found {
		//若已经包含, 则进行更新
		out := n.items[i]
		n.items[i] = item
		return out
	}
	//若没有找到, 进行插入逻辑的判断
	if len(n.children) == 0 {
		//若当前已经为叶子节点, 则直接加入叶子节点
		n.items.insertAt(i, item)
		return nil
	}
	//若没有找到且为中间节点, 则递归向孩子节点进行插入
	//在插入孩子节点之前需先判断孩子节点是否已经满了并进行分裂操作
	if n.maybeSplitChild(i, maxItems) {
		//孩子节点分裂完成, 有孩子的中间元素向本节点进行插入
		//先获取本节点的中间元素
		inTree := n.items[i]
		switch {
		case item.Less(inTree):
			//要查找的元素小于本节点的中间元素, 说明要向左(i)找孩子
		case inTree.Less(item):
			//要查找的元素大于于本节点的中间元素, 说明可以要向右(i+1)找孩子
			i++
		default:
			//若新加入的元素刚好是本节点的中间元素的话, 直接返回
			out := n.items[i]
			n.items[i] = item
			return out
		}
	}
	//向孩子节点进行查找
	return n.children[i].insert(item, maxItems)
}

// 进行分裂判断及分裂
func (n *node) maybeSplitChild(i, maxItems int) bool {
	if len(n.children[i].items) < maxItems {
		return false
	}
	first := n.children[i]
	item, second := first.split(maxItems / 2)
	//将孩子节点的中间元素插入当前节点
	n.items.insertAt(i, item)
	//将孩子节点新分裂出来的的节点加入当前节点的孩子节点范围
	n.children.insertAt(i+1, second)
	return true
}

//向index的位置插入节点
func (s *children) insertAt(index int, n *node) {
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = n
}

//向index的位置插入元素
func (s *items) insertAt(index int, item Item) {
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = item
}

// 使用二分查找, 若找到, 返回找到的索引, 若找不到, 返回的是中间位置的索引
func (s items) find(item Item) (index int, found bool) {
	i := sort.Search(len(s), func(i int) bool {
		return item.Less(s[i])
	})
	if i > 0 && !s[i-1].Less(item) {
		return i - 1, true
	}
	return i, false
}

// 分裂节点
func (n *node) split(i int) (Item, *node) {
	item := n.items[i]
	next := new(node)
	//将一半的元素分给新分裂出来的节点
	next.items = append(next.items, n.items[i+1:]...)
	//修剪当前节点分出去的元素
	n.items.truncate(i)
	if len(n.children) > 0 {
		//将一半的孩子节点分给新分裂的节点
		next.children = append(next.children, n.children[i+1:]...)
		//修剪当前节点分出去的孩子节点
		n.children.truncate(i + 1)
	}
	//返回中间的元素, 返回新分裂出来的节点
	return item, next
}

//修剪节点的元素
func (s *items) truncate(index int) {
	var toClear items
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, make(items, 16)):]
	}
}

//修剪节点的孩子节点
func (s *children) truncate(index int) {
	var toClear children
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, make(children, 16)):]
	}
}