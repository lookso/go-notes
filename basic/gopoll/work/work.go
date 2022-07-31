/*
@Time : 2019-10-20 19:55 
@Author : Tenlu
@File : work
@Software: GoLand
*/
package work

import "sync"

//任务类型接口
type Worker interface {
	Task(goid int)
}

// 任务池
type Pool struct {
	wg   sync.WaitGroup
	work chan Worker
}

func NewWorkPool(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	// 创建maxGoroutines个go协程
	for i := 0; i < maxGoroutines; i++ {
		go func(goid int) {
			//保证goroutine不停止执行通道中的任务
			for w := range p.work {
				w.Task(goid)
			}
			//每个goroutine不再执行work通道中任务时停止
			p.wg.Done()
		}(i)
	}
	return &p
}

// 运行
func (p *Pool) Run(r Worker) {
	p.work <- r
}

//停止
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
