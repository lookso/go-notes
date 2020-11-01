package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var rw = new(sync.RWMutex)
	var mt = make(map[int]string)
	var wg = new(sync.WaitGroup)
	wg.Add(2)
	go func(group *sync.WaitGroup) {
		defer group.Done()
		for i := 1; i <= 100; i++ {
			rw.RLock()
			fmt.Println(mt[i])
			rw.RUnlock()
		}
	}(wg)
	go func(group *sync.WaitGroup) {
		defer group.Done()
		for i := 1; i <= 100; i++ {
			rw.Lock()
			mt[i] = strconv.Itoa(i) + "-map-test"
			rw.Unlock()
		}
	}(wg)

	wg.Wait()

	syncMapFuc()
}

func syncMapFuc() {
	var sm sync.Map
	var wg = new(sync.WaitGroup)

	wg.Add(2)
	go func(group *sync.WaitGroup) {
		defer group.Done()
		for i := 1; i < 100; i++ {
			sm.Store(strconv.Itoa(i), i)
		}
	}(wg)
	go func(group *sync.WaitGroup) {
		defer group.Done()
		for i := 0; i < 100; i++ {
			v, _ := sm.Load(strconv.Itoa(i)) // 读
			fmt.Println(i, v)
		}
	}(wg)
	wg.Wait()

	var scene sync.Map
	//将键值对保存到sync.Map
	scene.Store("1", "a")
	scene.Store("2", "b")
	scene.Store("3", "c")
	scene.Store("4", "d")
	scene.Store("london", "london-val")
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
	channelMap()
}

var dataCh map[int]int = make(map[int]int)
var chMap chan int = make(chan int)

func channelMap() {
	// 并发启动的协程数量
	max := 10000
	time1 := time.Now().UnixNano()
	for i := 0; i < max; i++ {
		go modifyByChan(i)
	}
	// 处理channel的服务
	chanServ(max)
	time2 := time.Now().UnixNano()
	fmt.Printf("data len=%d, time=%d\n", len(dataCh), (time2-time1)/1000000)
}
func modifyByChan(i int) {
	chMap <- i
}
// 专门处理chMap的服务程序
func chanServ(max int) {
	for {
		i := <- chMap
		dataCh[i] = i
		if len(dataCh) == max {
			return
		}
	}
}
