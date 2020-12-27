package main

import "fmt"

func main() {
	obj := Constructor2()
	obj.AppendTail(1)
	obj.AppendTail(2)
	obj.AppendTail(3)
	popValue := obj.DeleteHead()
	fmt.Println(popValue)
}

type CQueue2 struct {
	in  []int // in 栈：存数据
	out []int // out 栈：出数据
}

func Constructor2() CQueue2 {
	return CQueue2{}
}

func (this *CQueue2) AppendTail(value int) {
	// 直接入 int 栈
	this.in = append(this.in, value)
}

func (this *CQueue2) DeleteHead() int {
	// 两个栈都为空
	if len(this.out) == 0 && len(this.in) == 0 {
		return -1
	}

	// out 栈为空，依次弹出 in 栈栈顶元素，入 out 栈
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			lastIndex := len(this.in) - 1
			popValue := this.in[lastIndex]
			this.in = this.in[:lastIndex]
			this.out = append(this.out, popValue)
		}
	}

	fmt.Println(this.out)
	// 弹出 out 栈顶元素
	lastIndex := len(this.out) - 1
	popValue := this.out[lastIndex]
	this.out = this.out[:lastIndex]
	return popValue
}
