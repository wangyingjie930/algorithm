/**
  @author: wangyingjie
  @since: 2023/3/6
  @desc:
**/

package distributeLock

import (
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	"time"
)

func newLock(i int) *redsync.Mutex {
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
	mutex := rs.NewMutex("my-lock", redsync.SetExpiry(100*time.Second))

	return mutex
}
