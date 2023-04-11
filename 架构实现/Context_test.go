package 架构实现

import (
	"context"
	"testing"
)

func Test_addNum(t *testing.T) {
	type args struct {
		numP      *int32
		id        int
		deferFunc func()
	}
	var a int32 = 1
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{numP: &a, id: 1, deferFunc: func() {
			t.Logf("defer")
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addNum(tt.args.numP, tt.args.id, tt.args.deferFunc)
			t.Log(*tt.args.numP)
		})
	}
}

func Test_coordinateWithWaitGroup(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "case1"},
		{name: "case2"},
		{name: "case3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coordinateWithWaitGroup()
		})
	}
}

/**
使用context改造wait_group
*/
func Test_coordinateWithContext(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "case1"},
		{name: "case2"},
		{name: "case3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coordinateWithContext()
		})
	}
}

func Test_Context(t *testing.T) {

	node1, cancelFunc1 := context.WithCancel(context.Background())
	defer cancelFunc1()

	// 示例1。
	node2 := context.WithValue(node1, "node2 key", "node2 value")
	node3 := context.WithValue(node2, "node3 key", "node3 value")
	//node1(根节点)->node2 ("node2 key", "node2 value")->node3("node3 key", "node3 value")

	//输出：node2 value， 从node3向上查找，知道根节点
	t.Logf("%v", node3.Value("node2 key"))
	//输出：node3 value
	t.Logf("%v", node3.Value("node3 key"))
	//输出：nil
	t.Logf("%v", node3.Value("node1 key"))
}
