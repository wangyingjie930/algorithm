package 链表法解决冲突

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"math"
)

// Bucket 键值对, 记住他们是合在一起的
type Bucket struct {
	Key string
	Value interface{}
	Next *Bucket //解决哈希冲突
}

// HashMap 散列表
type HashMap struct {
	data []*Bucket //存储一批键值对, 记住是数组形式
	cap int //容量
	len int //当前保存个数
}

func NewHashMap(cap int) *HashMap {
	return &HashMap{cap: cap, len: cap, data: make([]*Bucket, cap, cap)}
}

//使用哈希函数获取字符串对应数组索引下标
func (h *HashMap) getIndex(key string) int {
	var Result []byte
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(key))
	Result = Sha1Inst.Sum([]byte(""))

	binBuf := bytes.NewBuffer(Result)
	var x int32
	binary.Read(binBuf, binary.BigEndian, &x)
	return int(math.Abs(float64(x))) % h.cap
}

// Put 添加
func (h *HashMap) Put(key string, value interface{}) {
	//转换字符串为对应的下标
	index := h.getIndex(key)
	bucket := &Bucket{Key: key, Value: value}
	if h.data[index] == nil {
		//查看对应下标有没有键值对, 没有则直接添加
		h.data[index] = bucket
	}else {
		cur := h.data[index]
		var pre *Bucket
		for ;cur != nil;cur = cur.Next {
			//遍历, 当前key等于遍历的key, 进行修改更新
			if cur.Key == key {
				cur.Value = value
				return
			}
			pre = cur
		}
		//走到这里表明没有更新的节点, 往尾端插入
		if pre != nil {
			pre.Next = bucket
		}
	}
}

// Get 获取
func (h *HashMap) Get(key string) interface{}{
	index := h.getIndex(key)
	if h.data[index] == nil {
		return nil
	}else {
		//遍历对应的链表
		for cur := h.data[index]; cur != nil; cur = cur.Next {
			if key == cur.Key {
				return cur.Value
			}
		}
		return nil
	}
}

//rehash
func (h *HashMap) rehash() {
	//扩充2倍
	newHash := NewHashMap(h.cap * 2)
	//遍历键值对
	for _, bucket := range h.data {
		//遍历冲突链表
		for bucket != nil {
			//加入新的哈希表
			newHash.Put(bucket.Key, bucket.Value)
			bucket = bucket.Next
		}
	}
	//更新
	*h = *newHash
}