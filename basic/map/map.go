package main

import "fmt"
// https://zhuanlan.zhihu.com/p/273666774
func main() {
	m := make(map[int]int)
	m = map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	counter := 0
	for k, v := range m {
		counter++
		m[len(m)+1] = 1
		fmt.Println(k, v, len(m))
	}
	fmt.Println(m)
}
