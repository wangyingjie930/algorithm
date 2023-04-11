/**
  @author: wangyingjie
  @since: 2023/3/11
  @desc: copy from chatgpt
  https://www.haohongfan.com/docs/gohandbook/sync-chapter/2021-05-10-sync-cond/
**/

package 阻塞队列实现

import (
	"sync"
)

type BlockingQueue struct {
	queue    []interface{}
	capacity int
	mutex    sync.Mutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
}

func NewBlockingQueue(capacity int) *BlockingQueue {
	bq := &BlockingQueue{
		queue:    make([]interface{}, 0, capacity),
		capacity: capacity,
	}
	bq.notEmpty = sync.NewCond(&bq.mutex)
	bq.notFull = sync.NewCond(&bq.mutex)
	return bq
}

func (bq *BlockingQueue) Put(item interface{}) {
	bq.mutex.Lock()
	defer bq.mutex.Unlock()

	//注意这里是循环: 由于 wait 函数被唤醒时，存在虚假唤醒等情况，导致唤醒后发现，条件依旧不成立
	for len(bq.queue) == bq.capacity {
		//1. 当条件满足, 生成线程已经阻塞了, 没法进行后续的操作 (已经进行解锁mutex)
		bq.notFull.Wait()
	}

	bq.queue = append(bq.queue, item)
	bq.notEmpty.Signal()
}

func (bq *BlockingQueue) Get() interface{} {
	//2. 消费者线程可以正常进入
	bq.mutex.Lock()
	defer bq.mutex.Unlock()

	for len(bq.queue) == 0 {
		bq.notEmpty.Wait()
	}

	item := bq.queue[0]
	bq.queue = bq.queue[1:]
	bq.notFull.Signal()

	return item
}
