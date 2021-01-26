/*
@Time : 2019/5/12 10:22 PM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"context"
	"time"
)

func main()  {
	// 创建 context
	// 这个函数返回一个空 context。这只能用于高等级（在 main 或顶级请求处理中）
	ctx1:=context.Background()

	//这个函数也是创建一个空 context。也只能用于高等级或当您不确定使用什么 context，
	//或函数以后会更新以便接收一个 context 。这意味您（或维护者）计划将来要添加 context 到函数。
	ctx2:=context.TODO()
	/*
	这是它开始变得有趣的地方。此函数创建从传入的父 context 派生的新 context。父 context 可以是后台 context 或传递给函数的 context。
	返回派生 context 和取消函数。只有创建它的函数才能调用取消函数来取消此 context。如果您愿意，可以传递取消函数，但是，强烈建议不要这样做。
	这可能导致取消函数的调用者没有意识到取消 context 的下游影响。可能存在源自此的其他 context，这可能导致程序以意外的方式运行。简而言之，永远不要传递取消函数。
	*/
	//ctx3 := context.WithValue(context.Background(), interface{}{"age":12}, "test")

	/*
	此函数返回其父项的派生 context，当截止日期超过或取消函数被调用时，该 context 将被取消。
	例如，您可以创建一个将在以后的某个时间自动取消的 context，并在子函数中传递它。当因为截止日期耗尽而取消该 context 时，
	获此 context 的所有函数都会收到通知去停止运行并返回。
	*/
	ctx4, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))

	//此函数类似于 context.Wi
	// thDeadline。不同之处在于它将持续时间作为参数输入而不是时间对象。
	//此函数返回派生 context，如果调用取消函数或超出超时持续时间，则会取消该派生 context。
	ctx5, cancel := context.WithTimeout(context.Background(), 2 * time.Second)



}

