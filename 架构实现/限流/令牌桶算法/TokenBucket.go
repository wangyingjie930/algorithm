/**
  @author: wangyingjie
  @since: 2023/4/15
  @desc: https://ihui.ink/post/systemarch/00-limiter-algo/#4-%E4%BB%A4%E7%89%8C%E6%A1%B6%E7%AE%97%E6%B3%95
**/

package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	lock sync.Mutex

	rate     time.Duration // 多长时间生成一个令牌
	capacity int           // 桶的容量
	tokens   int           // 桶中当前token数量
	last     time.Time     // 桶上次放token的时间戳 s
}

func NewTokenBucket(bucketSize int, tokenRate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity: bucketSize,
		rate:     tokenRate,
	}
}

// 验证是否能获取一个令牌 返回是否被限流
func (t *TokenBucket) Allow() bool {
	t.lock.Lock()
	defer t.lock.Unlock()

	now := time.Now()
	if t.last.IsZero() {
		// 第一次访问初始化为最大令牌数
		t.last, t.tokens = now, t.capacity
	} else {
		if t.last.Add(t.rate).Before(now) {
			// 如果 now 与上次请求的间隔超过了 token rate
			// 则增加令牌，更新last
			t.tokens += int(now.Sub(t.last) / t.rate) // 将时间折算成token的数量 = (now - last) / rate
			if t.tokens > t.capacity {
				t.tokens = t.capacity
			}
			t.last = now
		}
	}

	if t.tokens > 0 {
		// 如果令牌数大于0，取走一个令牌
		t.tokens--
		return true
	}

	// 没有令牌,则拒绝
	return false
}

func main() {
	tokenBucket := NewTokenBucket(5, 100*time.Millisecond)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(limit *TokenBucket, i int) {
			defer wg.Done()
			fmt.Printf("第%d次请求: %t\n", i, tokenBucket.Allow())
		}(tokenBucket, i)
	}
	wg.Wait()
}
