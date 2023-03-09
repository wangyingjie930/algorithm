/**
  @author: wangyingjie
  @since: 2023/3/6
  @desc:
**/

package 分布式锁

import (
	"fmt"
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	"time"
)

func newLock(i int) {
	// 创建一个连接池
	pool := &redis.Pool{
		MaxIdle: 3,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}

	// 创建一个redsync对象
	rs := redsync.New([]redsync.Pool{pool})

	// 创建一个分布式互斥锁
	mutex := rs.NewMutex("my-lock", redsync.SetExpiry(3*time.Second))

	mutex.Lock()

	// 执行业务逻辑
	fmt.Print("proceeding: ", i, "\n")

	// 释放锁，如果失败则返回错误
	_, err := mutex.Unlock()
	if err != nil {
		panic(err)
	}
}
