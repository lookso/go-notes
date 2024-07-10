package main

import (
	"fmt"
	"sync"
)

// Result 结构体，用于存放处理结果和对应的索引
type Result struct {
	Index int
	Data  string // 假设处理结果是一个字符串
}

func main() {
	txtArr := []string{"text1", "text2", "text3"} // 待处理的文本数组
	resultsChan := make(chan Result, len(txtArr)) // 结果通道
	var wg sync.WaitGroup                         // 用于等待所有goroutine完成

	for i, txt := range txtArr {
		wg.Add(1)
		go func(i int, txt string) {
			defer wg.Done()
			// 模拟文本处理逻辑
			result := processText(i, txt)
			resultsChan <- result // 将结果发送到通道
		}(i, txt)
	}

	// 启动一个goroutine来关闭resultsChan通道
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// 从通道接收结果并存储
	resArr := make([]Result, len(txtArr))
	for res := range resultsChan {
		fmt.Println("res.Index", res.Index)
		resArr[res.Index] = res
	}

	// 输出结果
	for _, res := range resArr {
		fmt.Printf("Result %d: %s\n", res.Index, res.Data)
	}
}

// processText 模拟文本处理函数
func processText(index int, txt string) Result {
	// 这里只是简单地将文本转换为大写，实际逻辑可以更复杂
	return Result{Index: index, Data: txt + "_processed"}
}
