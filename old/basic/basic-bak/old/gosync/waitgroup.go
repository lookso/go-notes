/*
@Time : 2019-02-23 16:55 
@Author : Tenlu
@File : waitgroup
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)
// 同步的goroutine
// 由于goroutine是异步执行的，那很有可能出现主程序退出时还有goroutine没有执行完，此时goroutine也会跟着退出。
// 此时如果想等到所有goroutine任务执行完毕才退出,go提供了sync包和channel来解决同步问题，
// 当然如果你能预测每个goroutine执行的时间,你还可以通过time.Sleep方式等待所有的groutine执行完成以后在退出程序(如上面的列子)。

func cal(a int, b int, n *sync.WaitGroup) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	defer n.Done() //goroutinue完成后, WaitGroup的计数-1}
}

// WaitGroup 等待一组goroutinue执行完毕. 主程序调用 Add 添加等待的goroutinue数量.
// 每个goroutinue在执行结束时调用 Done ，此时等待队列数量减1.，主程序通过Wait阻塞，直到等待队列为0.
func main() {
	var goSync sync.WaitGroup // 声明一个WaitGroup变量
	for i := 0; i < 10; i++ {
		goSync.Add(1) // WaitGroup的计数加1
		go cal(i, i+1, &goSync)
	}
	goSync.Wait() //等待所有goroutine执行完毕}

	// 结果
	//4 + 5 = 9
	//1 + 2 = 3
	//9 + 10 = 19
	//5 + 6 = 11
	//2 + 3 = 5
	//6 + 7 = 13
	//7 + 8 = 15
	//8 + 9 = 17
	//3 + 4 = 7
	//0 + 1 = 1
}
