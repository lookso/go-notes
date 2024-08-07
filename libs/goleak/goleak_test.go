package goleak

import (
	"context"
	"time"

	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestGoroutineLeak(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// 模拟一些工作
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	neverClosed := make(chan bool)

	go func() {
		<-neverClosed // 这个协程将永远等待
	}()
	// go func() {
	// 	neverClosed <- true
	// }()

	// 给 Goroutine 一些时间运行
	time.Sleep(1 * time.Second)

	// 检查是否有泄露的 Goroutine
	if leak := goleak.Find(); leak != nil {
		t.Errorf("Found potential goroutine leak: %v", leak)
	}
}
