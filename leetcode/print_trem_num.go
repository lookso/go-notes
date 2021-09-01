package main

import (
	"fmt"
	"github.com/spf13/cast"
	"math"
	"runtime"
	"strconv"
	"sync"
)

// 0-2   1
// 3-5   2
// 6-8   3
// 9-11  4

// 一组数字,按照每5个元素长度起一个协程,不足5个的按5个算,有序打印出所有元素
func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	//var ch = make(chan int)
	var num = 5
	lf, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(len(arr))/float64(num)), 64)
	l := math.Ceil(lf)
	ll := cast.ToInt(l)
	fmt.Println("ll", ll)
	lock := sync.Mutex{}
	var stuIds = make([]int, 0)
	var wg = &sync.WaitGroup{}
	for i := 1; i <= ll; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("--" + strconv.Itoa(i) + "---")
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("panic err", err)
				}
				wg.Done()
			}()
			dl := i*num - 1
			if i > len(arr)/num {
				dl = len(arr) - 1
			}
			lock.Lock()
			for j := (i - 1) * num; j <= dl; j++ {
				stuIds = append(stuIds, arr[j]) // 并发写入确实有问题
			}
			lock.Unlock()
			fmt.Println("goroutine num", runtime.NumGoroutine())
		}(i)
	}
	wg.Wait()
	fmt.Println("goroutine num", runtime.NumGoroutine())
	fmt.Println("stu_ids", len(stuIds), stuIds)
}
