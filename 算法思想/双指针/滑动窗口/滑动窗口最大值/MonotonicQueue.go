/**
  @author: wangyingjie
  @since: 2023/5/5
  @desc: 单调队列的实现
**/

package 滑动窗口最大值

type MonotonicQueue struct {
	data []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{}
}

// Push
//  @Description: 队列中需要维护单调递减, 这也是核心
//  @receiver m
//  @param n
func (m *MonotonicQueue) Push(n int) {
	for len(m.data) > 0 && m.data[len(m.data)-1] < n {
		m.data = m.data[:len(m.data)-1]
	}
	m.data = append(m.data, n)
}

// Max
//  @Description: 直接取头一个即可
//  @receiver m
//  @return int
func (m *MonotonicQueue) Max() int {
	if len(m.data) > 0 {
		return m.data[0]
	}
	return -1
}

func (m *MonotonicQueue) Pop() {
	if len(m.data) > 0 {
		m.data = m.data[1:]
	}
}
