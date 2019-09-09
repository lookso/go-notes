/*
@Time : 2019-08-18 16:45 
@Author : Tenlu
@File : context
@Software: GoLand
*/
package main

import (
	"context"
	"fmt"
	"time"
)

var jobContextChan chan int
var ctx context.Context
var cancel context.CancelFunc

func contextWorker(jobContextChan <-chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case job := <-jobContextChan:
			fmt.Printf("执行任务 %d \n", job)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	jobContextChan = make(chan int, 100)
	//带有取消功能的 contex
	ctx, cancel = context.WithCancel(context.Background())
	//入队
	for i := 1; i <= 10; i++ {
		jobContextChan <- i
	}
	close(jobContextChan)

	go contextWorker(jobContextChan, ctx)
	time.Sleep(5 * time.Second)
	//調用cancel
	cancel()
}
