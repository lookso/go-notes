package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/top-k-frequent-elements/
func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 8, 8, 8, 6, 6, 6,6}
	k := 2
	fmt.Println(topKFrequent(nums, k))
	fmt.Println(topKFrequent2(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	// []int{1, 1, 1, 2, 2, 3,8,8,8,6,6,6}
	m := make(map[int]int)
	s := make([]int, 0)
	for _, v := range nums {
		i, ok := m[v]
		if ok {
			m[v] = i + 1
		} else {
			m[v] = 1
			s = append(s, v)
		}
	}
	fmt.Println("s", s)
	fmt.Println("m", m)

	sort.Slice(s, func(i, j int) bool {
		return m[s[i]] > m[s[j]]
	})
	fmt.Println("ss", s)
	fmt.Println(nums)

	return s[:k]
}
//	//作者：worker-boy
//	//链接：https://leetcode-cn.com/problems/top-k-frequent-elements/solution/xian-yong-mapcun-chu-xia-mei-ge-yuan-su-de-chu-xia/


func topKFrequent2(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/top-k-frequent-elements/solution/qian-k-ge-gao-pin-yuan-su-by-leetcode-solution/

