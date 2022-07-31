package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(mergeSection([][]int{{1, 3}, {0, 6}, {8, 10}, {15, 18}}))
}

func mergeSection(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println("intervals", intervals)
	res := make([][]int, 0)
	prev := intervals[0]
	fmt.Println("prev", prev)
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 没有一点重合
			res = append(res, prev)
			prev = cur
		} else { // 有重合
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//作者：xiao_ben_zhu
//链接：https://leetcode-cn.com/problems/merge-intervals/solution/shou-hua-tu-jie-56he-bing-qu-jian-by-xiao_ben_zhu/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
