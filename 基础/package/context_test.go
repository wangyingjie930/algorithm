package go_package

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("done")
	}

	return
}

func TestCancelContextMul(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		fmt.Println("开启子协程")
		go func(ctx context.Context) {
			fmt.Println("开启孙子协程")
			select {
			case <-ctx.Done():
				fmt.Println("孙子协程关闭")
				return
			}
		}(ctx)

		select {
		case <-ctx.Done():
			fmt.Println("子协程关闭")
			return
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(1 * time.Second) //等待协程的资源都关闭
}

func TestDeadlineContext(t *testing.T) {
	d := time.Now().Add(1 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// TestTimeOutContext
//  @Description: 和deadline一样的
//  @param t
func TestTimeOutContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel() //这个还是必要的, 可以从上级ctx的孩子ctx数组移除一个孩子

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}

// TestCancelContextWithValueContext
//  @Description: withvalue和其他的不一样, 没有返回cancel函数, 但是仍然可以和其他类型搭配使用
//  @param t
func TestCancelContextWithValueContext(t *testing.T) {
	cancelCtx, cancel := context.WithCancel(context.Background())
	ctx := context.WithValue(cancelCtx, "key1", "value1")
	ctx = context.WithValue(ctx, "key2", "value2")

	defer func() {
		cancel()
		time.Sleep(1 * time.Second)
	}()

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			t.Logf("finish: k1: %s", ctx.Value("key1"))
			t.Logf("finish: k2: %s", ctx.Value("key2"))
			return
		}
	}(ctx)
}
