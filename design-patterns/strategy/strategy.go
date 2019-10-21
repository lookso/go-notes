/*
@Time : 2019-10-21 08:18 
@Author : Tenlu
@File : strategy
@Software: GoLand
*/

//意图：定义一系列的算法,把它们一个个封装起来, 并且使它们可相互替换。
//主要解决：在有多种算法相似的情况下，使用 if...else 所带来的复杂和难以维护。

package main

import "fmt"

type ICompute interface {
	Compute(a, b int) int
}
type Add struct{}

func (p *Add) Compute(a, b int) int {
	return a + b
}

type Sub struct{}

func (p *Sub) Compute(a, b int) int {
	return a - b
}

// 重点在下面
var compute ICompute

type Context struct {
	A, B int
}

func NewContext(A, B int) *Context {
	return &Context{A: A, B: B}
}

func (p *Context) SetContext(o ICompute) {
	compute = o
}

func (p *Context) Result() int {
	return compute.Compute(p.A, p.B)
}

func main() {
	c := NewContext(15, 7)
	// 用户自己决定使用什么策略
	c.SetContext(new(Add))
	fmt.Println(c.Result())
}
