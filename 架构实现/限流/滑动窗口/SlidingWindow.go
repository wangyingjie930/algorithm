/**
  @author: wangyingjie
  @since: 2023/4/15
  @desc: https://ihui.ink/post/systemarch/00-limiter-algo/#%E9%80%82%E7%94%A8%E5%9C%BA%E6%99%AF
**/

package main

import (
	"fmt"
	"sync"
	"time"
)

type timeSlot struct {
	timestamp time.Time // 这个timeSlot的时间起点
	count     int       // 落在这个timeSlot内的请求数
}

func countReq(win []*timeSlot) int {
	var count int
	for _, ts := range win {
		count += ts.count
	}
	return count
}

type SlidingWindowLimiter struct {
	SlotDuration time.Duration // time slot的长度
	WinDuration  time.Duration // sliding window的长度
	numSlots     int           // window内最多有多少个slot
	maxReq       int           // win duration内允许的最大请求数
	windows      map[string][]*timeSlot
	sync.RWMutex
}

func NewSliding(slotDuration time.Duration, winDuration time.Duration, maxReq int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		SlotDuration: slotDuration,
		WinDuration:  winDuration,
		numSlots:     int(winDuration / slotDuration),
		windows:      make(map[string][]*timeSlot),
		maxReq:       maxReq,
	}
}

// 获取user_id/ip的时间窗口
func (l *SlidingWindowLimiter) getWindow(uidOrIp string) []*timeSlot {
	l.RLock()
	win, ok := l.windows[uidOrIp]
	l.RUnlock()

	if !ok {
		win = make([]*timeSlot, 0, l.numSlots)
	}
	return win
}

// storeWindow
//  @Description: map默认不支持并发安全
//  @receiver l
//  @param uidOrIp
//  @param win
func (l *SlidingWindowLimiter) storeWindow(uidOrIp string, win []*timeSlot) {
	l.Lock()
	l.windows[uidOrIp] = win
	l.Unlock()
}

// validate
//  @Description: 主要看这块
//  @receiver l
//  @param uidOrIp
//  @return bool
func (l *SlidingWindowLimiter) validate(uidOrIp string) bool {
	win := l.getWindow(uidOrIp)
	now := time.Now()

	// 循环遍历对应的slot, 直到最近的slot时间+window的长度=当前的时间
	timeoutOffset := -1
	for i, ts := range win {
		if ts.timestamp.Add(l.WinDuration).After(now) {
			break
		}
		timeoutOffset = i
	}
	if timeoutOffset > -1 {
		win = win[timeoutOffset+1:]
	}

	//计算这个window长度的slots的所有请求量是多少
	var result bool
	if countReq(win) < l.maxReq {
		result = true
	}

	// 记录这次的请求数
	var lastSlot *timeSlot
	if len(win) > 0 {
		lastSlot = win[len(win)-1]
		if lastSlot.timestamp.Add(l.SlotDuration).Before(now) {
			lastSlot = &timeSlot{timestamp: now, count: 1}
			win = append(win, lastSlot)
		} else {
			//注意点: 当前时间若处于最近的slot的时间范围内, 只需加1
			lastSlot.count++
		}
	} else {
		lastSlot = &timeSlot{timestamp: now, count: 1}
		win = append(win, lastSlot)
	}

	l.storeWindow(uidOrIp, win)

	return result
}

func (l *SlidingWindowLimiter) getUidOrIp() string {
	return "127.0.0.1"
}

func (l *SlidingWindowLimiter) IsLimited() bool {
	return !l.validate(l.getUidOrIp())
}

func main() {
	limiter := NewSliding(100*time.Millisecond, time.Second, 10)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(limit *SlidingWindowLimiter, i int) {
			defer wg.Done()
			fmt.Printf("第%d次请求: %t\n", i, limiter.IsLimited())
		}(limiter, i)
	}
	wg.Wait()
}
