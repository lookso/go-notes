/*
tts接口qps是200,
该接口有2个业务调用方c和B,c的服务请求量小,但是服务比较重要,尽量不失败,
b服务请求量大,但是可以重试,有哪些方案需要考虑？尽量保证接口 qps不被浪费也不被限制

拆分步骤
1.优先级队列 container/heap
2.根据优先级排序,请求数据去处理接口请求,请求完毕塞进channel
3.消费channel,返回结果

3.qps统计-缓存
4.业务
5.动态调整qps限制
*/
package main

import (
	"container/heap"
	"fmt"
)

// 优先级队列的元素，优先级越高，数值越小
type Request struct {
	Priority int // 优先级，数值越小优先级越高
	Value    string
	Index    int // 用于在heap中定位元素
}

// 优先级队列
type RequestQueue []*Request

// 实现heap.Interface接口
func (rq RequestQueue) Len() int           { return len(rq) }
func (rq RequestQueue) Less(i, j int) bool { return rq[i].Priority > rq[j].Priority }
func (rq RequestQueue) Swap(i, j int)      { rq[i], rq[j] = rq[j], rq[i] }

// 向优先级队列中添加元素
func (rq *RequestQueue) Push(x interface{}) {
	n := len(*rq)
	item := x.(*Request)
	item.Index = n
	*rq = append(*rq, item)
}

// 从优先级队列中移除并返回最高优先级的元素
func (rq *RequestQueue) Pop() interface{} {
	old := *rq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // 将Index标记为无效
	*rq = old[0 : n-1]
	return item
}

func main() {
	rq := &RequestQueue{}
	heap.Init(rq)
	for i := 0; i < 100; i++ {
		priority := 50
		request := &Request{
			Priority: priority,
			Value:    fmt.Sprintf("B-Request %d", i),
		}
		heap.Push(rq, request)
	}
	for i := 0; i < 10; i++ {
		priority := 50
		request := &Request{
			Priority: priority,
			Value:    fmt.Sprintf("BB-Request %d", i),
		}
		heap.Push(rq, request)
	}
	for i := 0; i < 10; i++ {
		priority := 100 // C端的请求优先级更高
		request := &Request{
			Priority: priority,
			Value:    fmt.Sprintf("C-Request %d", i),
		}
		heap.Push(rq, request)
	}

	// 处理所有请求
	for rq.Len() > 0 {
		request := heap.Pop(rq).(*Request)
		fmt.Printf("Handling request with priority %d: %s\n", request.Priority, request.Value)
	}
}
