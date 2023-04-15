/**
  @author: wangyingjie
  @since: 2023/3/11
  @desc:
**/

package BlockingQueue

import (
	"fmt"
	"testing"
	"time"
)

func TestNewBlockingQueue(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *BlockingQueue
	}{
		{name: "case1", args: args{capacity: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewBlockingQueue(tt.args.capacity)
			go func(queue *BlockingQueue) {
				//消费
				for true {
					time.Sleep(1000 * time.Millisecond)
					fmt.Println("consume: ", queue.Get())
				}
			}(queue)

			for i := 0; i < 100; i++ {
				queue.Put(i)
				fmt.Println("product: ", i)
			}
		})
	}
}
