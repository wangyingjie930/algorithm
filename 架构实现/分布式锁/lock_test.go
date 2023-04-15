/**
  @author: wangyingjie
  @since: 2023/3/6
  @desc:
**/

package distributeLock

import (
	"github.com/go-redsync/redsync"
	"sync"
	"testing"
)

func Test_newLock(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{i: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lock := newLock(tt.args.i)

			var wg sync.WaitGroup
			for i := 0; i < 20; i++ {
				wg.Add(1)
				go func(lock *redsync.Mutex, i int) {
					lock.Lock()
					defer func() {
						lock.Unlock()
						wg.Done()
					}()

					t.Log("i:", i)
				}(lock, i)
			}

			wg.Wait()
		})
	}
}
