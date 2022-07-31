package main

import (
	"fmt"
	"sync"
	"time"
)

type IsOnlineItem struct {
	WxUserID string `json:"wx_user_id"`
	Online   int    `json:"online"`
}

func useWaitGroupBug() {
	var wg = new(sync.WaitGroup)
	lock := sync.Mutex{}
	list := make([]IsOnlineItem, 0)
	wxUserIDs := []string{"01", "02", "03"}
	for _, wxUserID := range wxUserIDs {
		wg.Add(1)
		go func(wxUserID string) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("IsOnline panic err:%+v", err)
				}
				wg.Done()
			}()
			lock.Lock()
			list = append(list, IsOnlineItem{
				WxUserID: wxUserID,
				Online:   1,
			})
			lock.Unlock()
		}(wxUserID)
	}
	wg.Wait()
	fmt.Printf("%v\n", list)
}
func main() {
	useWaitGroupBug()
	num := 3
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

// 有一点需要特别注意的是process()中使用指针类型的*sync.WaitGroup作为参数，
// 这里不能使用值类型的sync.WaitGroup作为参数，因为这意味着每个goroutine都拷贝一份wg，每个goroutine都使用自己的wg。
// 这显然是不合理的，这3个goroutine应该共享一个wg，才能知道这3个goroutine都完成了。实际上，如果使用值类型的参数，main goroutine将会永久阻塞而导致产生死锁。
