/**
  @author: wangyingjie
  @since: 2023/4/11
  @desc: 设计熔断器 copy from chatgpt
**/

package breaker

import (
	"fmt"
	"math/rand"
	"time"
)

type CircuitBreaker struct {
	State                   State
	TripThreshold           int           `comment:"表示触发熔断器开启的阈值，即在多少次连续错误之后触发熔断器开启"`
	SuccessThreshold        int           `comment:"表示在半开启状态下需要达到的连续成功次数，如果达到了就将熔断器的状态设置为关闭状态"`
	ConsecutiveSuccessCount int           // 表示连续成功的次数，在判断是否需要将熔断器从半开启状态转换为关闭状态时使用
	ConsecutiveErrorCount   int           // 表示连续错误的次数，在判断是否需要触发熔断器开启、或者将熔断器从半开启状态转换为开启状态时使用
	SleepWindow             time.Duration // 表示在熔断器开启后需要等待多长时间才能尝试再次调用服务，这可以避免熔断器过于频繁地尝试调用服务，进一步减轻服务端的负担
	LastAttemptTime         time.Time     //表示上一次服务调用的时间，用于计算是否已经过了等待时间（SleepWindow
}

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

func (cb *CircuitBreaker) Call() error {
	switch cb.State {
	case Closed:
		if cb.ConsecutiveErrorCount >= cb.TripThreshold {
			fmt.Println("Tripping circuit breaker to open state")
			cb.State = Open
			cb.LastAttemptTime = time.Now()
			return fmt.Errorf("circuit breaker is open")
		}
		err := cb.doCall()
		if err != nil {
			fmt.Printf("Call failed: %v\n", err)
			cb.ConsecutiveSuccessCount = 0
			cb.ConsecutiveErrorCount++
		} else {
			fmt.Println("Call succeeded")
			cb.ConsecutiveSuccessCount++
			cb.ConsecutiveErrorCount = 0
		}
		return err

	case Open:
		if time.Since(cb.LastAttemptTime) >= cb.SleepWindow {
			fmt.Println("Trying again...")
			cb.State = HalfOpen
			cb.LastAttemptTime = time.Now()
			return cb.Call()
		}
		fmt.Println("Circuit breaker is open")
		return fmt.Errorf("circuit breaker is open")

	case HalfOpen:
		err := cb.doCall()
		if err != nil {
			fmt.Printf("Call failed: %v\n", err)
			cb.State = Open
		} else {
			fmt.Println("Call succeeded")
			cb.State = Closed
		}
		return err

	default:
		return fmt.Errorf("unknown state")
	}
}

func (cb *CircuitBreaker) doCall() error {
	fmt.Println("Making call...")
	// 模拟服务调用的随机失败率
	if rand.Float32() < 0.2 {
		return fmt.Errorf("random failure")
	}
	return nil
}
