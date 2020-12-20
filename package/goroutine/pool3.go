package main

import (
	. "go-notes/package/pool"
	"log"
	"sync"
	"time"
)

//
var names = []string{
	"lili",
	"yingying",
}

//Worker实现类型
type namePrinter struct {
	name string
}

func (n *namePrinter) Task(goid int) {
	log.Printf("goroutineID:%d，打印名字为：%s\n", goid, n.name)
	time.Sleep(time.Second)
}

func main() {
	p := New(3)
	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	for i := 0; i < 10; i++ {
		for _, name := range names {
			//任务实例
			np := namePrinter{
				name: name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
