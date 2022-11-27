package 链表法解决冲突

import "testing"

func Test_getIndex(t *testing.T) {
	h := NewHashMap(8)
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{key: "test1"}},
		{name: "case1", args: args{key: "test2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := h.getIndex(tt.args.key)
			t.Logf("getIndex() = %v", got)
		})
	}
}

func TestHashMap_Put(t *testing.T) {
	h := NewHashMap(10)
	h.Put("test1", 1)
	h.Put("test2", 2)
	t.Log(h.Get("test1"), h.getIndex("test1"), h.Get("test2"), h.getIndex("test2"))

	h.Put("test1", 3)
	h.Put("test2", 6)
	t.Log(h.Get("test1"), h.getIndex("test1"), h.Get("test2"), h.getIndex("test2"))

	h.rehash()
	t.Log(h.Get("test1"), h.getIndex("test1"), h.Get("test2"), h.getIndex("test2"))
}