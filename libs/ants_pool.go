package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	// 创建一个容量为10的协程池
	p, _ := ants.NewPool(10)
	defer p.Release()

	// 使用WaitGroup等待所有任务完成
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		j := i
		wg.Add(1)
		// 提交任务到协程池
		_ = p.Submit(func() {
			//	defer wg.Done()
			// 模拟任务执行
			time.Sleep(2 * time.Second)
			fmt.Printf("任务" + strconv.Itoa(j) + "执行完成\n")
			fmt.Printf("run times: %d, waiting times: %d\n", p.Running(), p.Waiting())
		})
	}

	// 等待所有任务完成
	wg.Wait()
	fmt.Println("-------------------")
}
