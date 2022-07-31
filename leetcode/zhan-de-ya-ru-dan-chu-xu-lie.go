package main

import "fmt"

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, popped))
}
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	res := false
	// 模拟进栈
	for i := 0; i < len(pushed); i++ {
		// 进栈
		stack = append(stack, pushed[i])
		// 出栈
		for popped[0] == stack[len(stack)-1] {
			//fmt.Println("popped",popped)
			popped = popped[1:]
			stack = stack[:len(stack)-1]
			// 防止slice越界
			if len(stack) == 0 {
				break
			}
		}
	}
	if len(popped) == 0 {
		res = true
	}
	return res
}

//https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/solution/jian-zhi-offer-31-zhan-de-ya-ru-dan-chu-xu-lie-mo-/
