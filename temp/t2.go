package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 定义数据结构，包含唯一ID和其他数据
type Data struct {
	ID    string
	Value int
}

func main() {
	for i := 0; i < 20; i++ {

		// 创建一个用于发送数据的通道
		dataChan := make(chan Data)

		// 生成一个唯一ID
		uniqueID := generateUniqueID()

		// 启动一个goroutine来处理数据发送
		var wg sync.WaitGroup
		wg.Add(1)
		go func(id string, ch chan<- Data, wg *sync.WaitGroup) {
			defer wg.Done()
			// 模拟数据生成和发送

			// 生成包含唯一ID的数据
			data := Data{
				ID:    id,
				Value: rand.Intn(100),
			}
			// 发送数据到通道
			ch <- data
			// 模拟耗时操作
			//			time.Sleep(time.Second)

			// 关闭通道（如果不需要再发送数据）
			close(ch)
		}(uniqueID, dataChan, &wg)

		// 从通道接收数据并处理
		for data := range dataChan {
			fmt.Printf("k:%+v,uniqueID:%+v,Received data: ID=%s, Value=%d\n", i, uniqueID, data.ID, data.Value)
		}

		// 等待goroutine完成
		wg.Wait()
	}
}

// 生成唯一ID的函数（简单示例，实际中可能需要更复杂的逻辑）
func generateUniqueID() string {
	return fmt.Sprintf("ID-%d", rand.Int())
}
