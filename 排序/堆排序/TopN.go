package Heap

import "fmt"

	/**
	topN算法
	 */
	type TopN struct {
		data []int
		count int
		N int
	}

	func NewTopN(count int) TopN {
		data := []int{0}
		return TopN{data: data, count: 0, N: count}
	}

	func (t *TopN) Insert (item int) {
		//元素未满直接插入, 构建大顶堆
		if t.count < t.N {
			t.data = append(t.data, item)
			t.count ++
			t.shiftUp(t.count)
		}else {
			//如果当前元素比堆顶元素小, 替换堆顶元素, 执行shiftdown重新构建大顶堆
			if t.data[1] > item {
				t.data[1] = item
				t.shiftDown(1)
			}
		}
	}

	func (t *TopN) Data()  {
		fmt.Println(t.data)
	}

	/**
	插入完之后重新排列, 使它符合大顶堆条件
	*/
	func (t *TopN) shiftUp(i int) {
		for ;i > 1 && t.data[i] > t.data[i / 2]; i = i / 2 {
			t.data[i], t.data[i / 2] = t.data[i / 2], t.data[i]
		}
	}

	func (t *TopN) shiftDown(i int)  {
		for ;2 * i <= t.count; {
			//默认取左节点
			j := 2 * i
			//左 < 右, 取右
			if j + 1 <= t.count && t.data[j] < t.data[j + 1] {
				j = j + 1
			}
			//当前节点大于左右节点, 跳出
			if t.data[i] >= t.data[j] {
				break
			}
			//当前节点与 左/右 节点交换
			t.data[i], t.data[j] = t.data[j], t.data[i]
			i = j //向前移动
		}
	}