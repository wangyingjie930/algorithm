/**
  @author: wangyingjie
  @since: 2023/4/15
  @desc: https://ihui.ink/post/systemarch/00-limiter-algo/#go%E5%AE%9E%E7%8E%B0-2
https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/

不太适合电商抢购和微博出现热点事件等场景的限流，
一是应对突发流量不是很灵活，
二是为每个user_id/ip维护一个队列(木桶)，worker从这些队列中拉取任务，资源的消耗会比较大
**/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 封装业务逻辑的执行结果
type Result struct {
	Msg string
}

// 执行的业务逻辑函数
type Handler func() Result

// 每个请求来了，把需要执行的业务逻辑封装成Task，放入木桶，等待worker取出执行
type Task struct {
	handler Handler     // worker从木桶中取出请求对象后要执行的业务逻辑函数
	resChan chan Result // 等待worker执行并返回结果的channel
	taskID  int
}

func NewTask(id int, handler Handler) Task {
	return Task{
		handler: handler,
		resChan: make(chan Result),
		taskID:  id,
	}
}

// 漏桶
type LeakyBucket struct {
	BucketSize int       // 木桶的大小
	WorkerNum  int       // 同时从木桶中获取任务执行的worker数量
	bucket     chan Task // 存方任务的木桶
}

func NewLeakyBucket(bucketSize int, workNum int) *LeakyBucket {
	return &LeakyBucket{
		BucketSize: bucketSize,
		WorkerNum:  workNum,
		bucket:     make(chan Task, bucketSize),
	}
}

func (b *LeakyBucket) AddTask(task Task) bool {
	// 如果木桶已经满了，返回false
	select {
	case b.bucket <- task:
	default:
		fmt.Printf("request[id=%d] is refused\n", task.taskID)
		return false
	}

	// 如果成功入桶，调用者会等待worker执行结果
	resp := <-task.resChan
	fmt.Printf("request[id=%d] is run ok, resp[%v]\n", task.taskID, resp)
	return true
}

// Start
//  @Description: 当主函数的defer调用到cancel时, 会触发ctx.Done, 此时协程结束
//  @receiver b
//  @param ctx
func (b *LeakyBucket) Start(ctx context.Context) {
	// 开启worker从木桶拉取任务执行
	for i := 0; i < b.WorkerNum; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					task := <-b.bucket
					result := task.handler()
					task.resChan <- result
				}
			}
		}(ctx)
	}
}

func main() {
	bucket := NewLeakyBucket(10, 4)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bucket.Start(ctx) // 开启消费者

	// 模拟20个并发请求
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func(id int) {
			defer wg.Done()
			task := NewTask(id, func() Result {
				time.Sleep(3000 * time.Millisecond)
				return Result{}
			})
			bucket.AddTask(task)
		}(i)
	}
	wg.Wait()
	time.Sleep(10 * time.Second)
}
