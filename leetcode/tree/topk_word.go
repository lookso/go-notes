package tree

import "container/heap"

type node struct {
	s   string
	num int
}
type nodes []node

//返回true是要保持的模样
func (n nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n nodes) Less(i, j int) bool {
	if n[i].num == n[j].num {
		return n[i].s < n[j].s
	}
	return n[i].num > n[j].num
}
func (n nodes) Len() int {
	return len(n)
}
func (n *nodes) Pop() interface{} {
	var t = (*n)[0]
	old := *n
	old[0] = old[n.Len()-1]
	old = old[:n.Len()-1]
	*n = old
	return t
}

func (n *nodes) Push(x interface{}) {
	*n = append(*n, x.(node))
}

func topKFrequent(words []string, k int) []string {
	var p = make(map[string]int)
	for _, v := range words {
		p[v]++
	}
	var n nodes
	for s, v := range p {
		heap.Push(&n, node{s: s, num: v})
	}
	heap.Init(&n)
	var res []string
	for k != 0 {
		var t = n.Pop().(node)
		heap.Fix(&n, 0)
		res = append(res, t.s)
		k--
	}
	return res
}

//
//作者：llwlzh
//链接：https://leetcode-cn.com/problems/top-k-frequent-words/solution/xie-tian-xie-di-ni-guo-liao-by-llwlzh/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
