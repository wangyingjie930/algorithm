/**
  @author: wangyingjie
  @since: 2023/3/11
  @desc: copy from chatgpt
  map: 多个节点/协程各自处理一部分数据
  reduce: 将多个节点/协程得到的结果合并
**/

package MapReduce

import (
	"strconv"
	"strings"
	"sync"
)

type KeyValue struct {
	Key   string
	Value string
}

func mapper(input string) []KeyValue {
	// 根据输入生成键值对列表
	// 这里简单地将每个单词作为键，值为1
	words := strings.Fields(input)
	var kvs []KeyValue
	for _, w := range words {
		kvs = append(kvs, KeyValue{w, "1"})
	}
	return kvs
}

func reducer(key string, values []string) string {
	// 对相同键的值进行统计并返回结果
	count := 0
	for _, v := range values {
		n, _ := strconv.Atoi(v)
		count += n
	}
	return strconv.Itoa(count)
}

func runMapReduce(input []string, mapper func(string) []KeyValue, reducer func(string, []string) string) map[string]string {
	// 创建一个WaitGroup，用于等待所有协程完成
	var wg sync.WaitGroup
	// 创建一个互斥锁，用于保护共享变量
	var mu sync.Mutex
	// 创建一个映射表，用于保存中间结果
	kvs := make(map[string][]string)

	// 对输入数据进行遍历，为每个输入数据创建一个协程
	for _, in := range input {
		// 增加WaitGroup计数器
		wg.Add(1)
		// 启动一个协程进行处理
		go func(in string) {
			// 在协程结束时减少WaitGroup计数器
			defer wg.Done()
			// 调用mapper函数将输入数据转换为中间结果
			intermediate := mapper(in)
			// 将中间结果添加到映射表中
			for _, kv := range intermediate {
				// 为了保证线程安全，需要使用互斥锁 (同一时间只能由一个协程操作映射表)
				mu.Lock()
				kvs[kv.Key] = append(kvs[kv.Key], kv.Value)
				mu.Unlock()
			}
		}(in)
	}
	// 等待所有协程完成
	wg.Wait()

	// 创建一个映射表，用于保存最终结果
	output := make(map[string]string)
	// 遍历中间结果，并调用reducer函数将相同键的值合并为一个结果
	for k, v := range kvs {
		output[k] = reducer(k, v)
	}
	// 返回最终结果
	return output
}

/*func startMapReduce() {
	input := []string{
		"hello world",
		"world world",
		"world hello",
		"go go go",
		"go world",
	}
	result := runMapReduce(input, mapper, reducer)
	for k, v := range result {
		fmt.Printf("%s: %s\n", k, v)
	}
}*/
