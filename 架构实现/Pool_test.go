package 架构实现

import (
	"testing"
)

func Test_initPool(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initPool()
			p := pool.Get().(*Person)
			t.Logf("首次从 pool 里获取：%v", p)

			p.Name = "first"
			t.Logf("设置 p.Name = %s\n", p.Name)
			pool.Put(p)
			t.Logf("get again %v", pool.Get().(*Person))
			t.Logf("get again %v", pool.Get().(*Person))

			//get完之后pool里面就没有数据了
			//但是并发场景无法保证这种顺序，最好的做法是在 Put 前，将对象清空
			pool.Put(p)
			t.Logf("get again %v", pool.Get().(*Person))
			t.Logf("get again %v", pool.Get().(*Person))
		})
	}
}
