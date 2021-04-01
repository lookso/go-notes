package main

import (
	"container/heap"
	"errors"
	"fmt"
)
// https://blog.csdn.net/hp_cpp/article/details/102538863

// TopK问题-基于堆排序和快速排序的实现
// https://studygolang.com/articles/25659

// 笔试面试算法经典-打印n个数组中最大的topk
// https://www.geek-share.com/detail/2703334141.html

type HeapNode struct {
	value  int32 //值是什么
	arrNum int   //来自哪个数组
	index  int   //来自数组的哪个位置
}

type ArrayHeap []*HeapNode

func (arrayheap ArrayHeap) Len() int {
	return len(arrayheap)
}

func (arrayheap ArrayHeap) Less(i, j int) bool {
	return arrayheap[i].value > arrayheap[j].value //大根堆
}

func (arrayheap ArrayHeap) Swap(i, j int) {
	arrayheap[i], arrayheap[j] = arrayheap[j], arrayheap[i]
}

func (arrayheap *ArrayHeap) Push(node interface{}) {
	// n := len(*arrayheap)
	item := node.(*HeapNode)
	*arrayheap = append(*arrayheap, item)
}

func (arrayheap *ArrayHeap) Pop() interface{} {
	old := *arrayheap
	n := len(old)
	item := old[n-1]
	*arrayheap = old[0 : n-1]
	return item
}

func (arrayheap *ArrayHeap) IsEmpty() bool {
	length := len(*arrayheap)
	return length == 0
}

func getMatriValue(matrix *[3][5]int32, i, j int) (value int32, err error) {
	if i < 0 || i > len(*matrix)-1 {
		return 0, errors.New("out of matrix index")
	}
	if j < 0 || j > len(matrix[0])-1 {
		return 0, errors.New("out of matrix index")
	}
	return matrix[i][j], nil
}

// func printTopK( int topK) {

// }

func main() {
	var matrix = [3][5]int32{
		{219, 405, 538, 845, 971},
		{148, 558, 0, 0, 0},
		{52, 99, 348, 691, 0},
	}

	myHeap := make(ArrayHeap, 0)
	heap.Init(&myHeap)

	heap.Push(&myHeap, &HeapNode{
		value:  971,
		arrNum: 0,
		index:  6,
	})

	heap.Push(&myHeap, &HeapNode{
		value:  558,
		arrNum: 1,
		index:  1,
	})

	heap.Push(&myHeap, &HeapNode{
		value:  691,
		arrNum: 2,
		index:  3,
	})

	for i := 0; i < 5; i++ {
		if bEmpty := myHeap.IsEmpty(); bEmpty {
			break
		}

		item := heap.Pop(&myHeap).(*HeapNode)
		fmt.Printf("value = %d, arrNum = %d, index = %d\n", item.value, item.arrNum, item.index)

		if nextValue, ok := getMatriValue(&matrix, item.arrNum, item.index-1); ok == nil {
			heap.Push(&myHeap, &HeapNode{
				value:  nextValue,
				arrNum: item.arrNum,
				index:  item.index - 1,
			})
			length := len(myHeap)
			heap.Fix(&myHeap, length-1)
		}
	}
}

//————————————————
//版权声明：本文为CSDN博主「hp_cpp」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
//原文链接：https://blog.csdn.net/hp_cpp/article/details/102538863
