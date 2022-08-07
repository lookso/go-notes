/*
@Time : 2019-10-21 09:39
@Author : Tenlu
@File : decorate
@Software: GoLand
*/
package main

import (
	"fmt"
	"math"
)

// 装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。
// 这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。

// 动态地给一个对象添加一些额外的职责。就增加功能来说，装饰器模式相比生成子类更为灵活。
// 装饰器。实现接口，又定义了自己的事件DecorateFun，相当于抽象类

//关键代码： 1、Component 类充当抽象角色，不应该具体实现。 2、修饰类引用和继承 Component 类，具体扩展类重写父类方法。
//应用实例： 1、孙悟空有 72 变，当他变成"庙宇"后，他的根本还是一只猴子，但是他又有了庙宇的功能

// 动态地给一个对象增加一些额外的职责，就拓展对象功能来说，装饰模式比生成子类的方式更为灵活。

type IDecorate interface {
	Do()
}

type Decorate struct {
	// 待装饰的抽象类
	decorate IDecorate
}

func (c *Decorate) DecorateFun(i IDecorate) {
	c.decorate = i
}

func (c *Decorate) Do() {
	if c.decorate != nil {
		c.decorate.Do()
	}
}

// 具体A装饰
type Person struct {
	Base Decorate
}

// 重写方法，隐式实现接口
func (c *Person) Do() {
	fmt.Printf("执行Person-A装饰\n")
	c.Base.Do()
}

// 具体B装饰
type Clothes struct {
	Base Decorate
}

// 重写方法，隐式实现接口
func (c *Clothes) Do() {
	fmt.Printf("执行clothes-B装饰\n")
	c.Base.Do()
}

func main() {
	// 添加装饰A
	p := new(Person)
	// 添加装饰B
	c := new(Clothes)
	p.Base.DecorateFun(c)
	// 执行
	p.Do()
	var num = "100"
	fmt.Println(myAtoi(num))

}
func myAtoi(str string) int {
	res, flag := 0, 1
	for _, s := range str {
		switch s {
		case ' ','+','-','.':
			return -1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			res = res*10 + int((s - '0'))
			if res*flag > math.MaxInt32 {
				return math.MaxInt32
			}
			if res*flag < math.MinInt32 {
				return math.MinInt32
			}
		default:
			return -1
		}
	}
	return res * flag
}
