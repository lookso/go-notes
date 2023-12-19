package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	// 创建一个带有超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 创建一个 errgroup
	g, ctx := errgroup.WithContext(ctx)

	g.SetLimit(2) // 设置并发限制为2

	for i := 0; i < 10; i++ {
		j := i // 创建循环变量的副本，以避免闭包引用问题
		// 启动一个会超时的任务
		g.Go(func() error {
			fmt.Println("started", j)
			select {
			case <-time.After(3 * time.Second):
				fmt.Println("finished")
				return nil
			case <-ctx.Done():
				// 如果 context 超时，就返回一个错误
				errors.New("timeout")
				return ctx.Err()
			}
		})
		fmt.Println("j", j)
	}

	if err := g.Wait(); err != nil {
		fmt.Println("encountered error:", err)
	} else {
		fmt.Println("all goroutines finished without error")
	}
}
