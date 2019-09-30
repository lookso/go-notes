package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
1.context.Background() Context
这个函数返回一个空 context。这只能用于高等级（在 main 或顶级请求处理中）。这能用于派生我们稍后谈及的其他 context 。

2.context.TODO() Context
这个函数也是创建一个空 context。也只能用于高等级或当您不确定使用什么 context,或函数以后会更新以便接收一个 context 。
这意味您（或维护者）计划将来要添加 context 到函数。

3.context.WithValue(parent Context, key, val interface{}) (ctx Context, cancel CancelFunc)
此函数接收 context 并返回派生 context，其中值 val 与 key 关联，并通过 context 树与 context 一起传递。
这意味着一旦获得带有值的 context，从中派生的任何 context 都会获得此值。
不建议使用 context 值传递关键参数，而是函数应接收签名中的那些值，使其显式化。
ctx := context.WithValue(context.Background(), key, "test")

4.context.WithCancel()
这是它开始变得有趣的地方。此函数创建从传入的父 context 派生的新 context。
父 context 可以是后台 context 或传递给函数的 context。
返回派生 context 和取消函数。只有创建它的函数才能调用取消函数来取消此 context。
如果您愿意，可以传递取消函数，但是，强烈建议不要这样做。这可能导致取消函数的调用者没有意识到取消 context 的下游影响。
可能存在源自此的其他 context，这可能导致程序以意外的方式运行。简而言之，永远不要传递取消函数。
ctx, cancel := context.WithCancel(context.Background())

5.context.WithTimeout()
ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
不同之处在于它将持续时间作为参数输入而不是时间对象。此函数返回派生 context，如果调用取消函数或超出超时持续时间，则会取消该派生 context。

通过上面我们知道了如何创建 context（Background 和 TODO）以及如何派生 context（WithValue，WithCancel，Deadline 和 Timeout），

6.下面可以看到接受 context 的函数启动一个 goroutine 并等待 该 goroutine 返回或该 context 取消。select 语句帮助我们选择先发生的任何情况并返回。
<-ctx.Done() 一旦 Done 通道被关闭，这个 <-ctx.Done(): 被选择。一旦发生这种情况，此函数应该放弃运行并准备返回。
这意味着您应该关闭所有打开的管道,释放资源并从函数返回。
有些情况下，释放资源可以阻止返回，比如做一些挂起的清理等等。在处理context 返回时,您应该注意任何这样的可能性。

超时控制
context.WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) 指定时长超时结束
context.WithCancel(parent Context) (ctx Context, cancel CancelFunc) 手动结束
context.WithDeadline(parent Context, d time.Time) (Context, CancelFunc) 指定时间结束

*/
func main() {
	//stopCtx, doStop := context.WithCancel(context.Background())
	stopCtx, doStop := context.WithTimeout(context.TODO(), time.Second*3)
	defer doStop() // 防止任务比超时时间短导致资源未释放
	// 启动协程
	go task(stopCtx)
	// 主协程需要等待，否则直接退出
	stopSignal(doStop) //平滑停止信号监听
	select {}
}

func task(ctx context.Context) {
	ch := make(chan struct{}, 0)
	// 真正的任务协程
	go func() {
		// 模拟两秒耗时任务
		time.Sleep(time.Second * 4)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("Goroutine Is Stop")
		return
	}
}

/**
 * 停止信号监听并平滑关闭服务
 * @params doStop	停止上下文关闭方法
 */
func stopSignal(doStop context.CancelFunc) {
	fmt.Println(os.Getpid())
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGTERM)
	go func(doStop context.CancelFunc) {
		sig := <-stopSignal
		if sig == syscall.SIGTERM {
			doStop()
		}
	}(doStop)
}
