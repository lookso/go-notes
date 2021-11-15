package main

// 建造者模式:将一个复杂对象的构造与它的表示分离，使同样的构建过程可以创建不同的表示。
import "fmt"

//定义一个产品(待构建的产品)
type People struct {
	Hat     string //帽子
	Clothes string //衣服
	Pants   string //裤子
	Sock    string //袜子
	Shoes   string //鞋子
}

//展示产品信息
func (p People) ShowPeopleInfo() {
	fmt.Println(p)
}

//建造者接口
type Builder interface {
	SetHat(hat string) Builder
	SetClothes(clothes string) Builder
	SetPants(pants string) Builder
	SetSock(sock string) Builder
	SetShoes(shoes string) Builder
	Build() *People
}

//具体的建造者
type ConcreteBuilder struct {
	//需要建造的产品
	people *People
}

func (c *ConcreteBuilder) SetHat(hat string) Builder {
	c.init()
	c.people.Hat = hat
	return c
}

func (c *ConcreteBuilder) SetClothes(clothes string) Builder {
	c.init()
	c.people.Clothes = clothes
	return c
}

func (c *ConcreteBuilder) SetPants(pants string) Builder {
	c.init()
	c.people.Pants = pants
	return c
}

func (c *ConcreteBuilder) SetSock(sock string) Builder {
	c.init()
	c.people.Sock = sock
	return c
}

func (c *ConcreteBuilder) SetShoes(shoes string) Builder {
	c.init()
	c.people.Shoes = shoes
	return c
}

func (c *ConcreteBuilder) Build() *People {
	return c.people
}

//初始化具体建造者的数据结构
func (c *ConcreteBuilder) init() {
	if c.people == nil {
		c.people = &People{}
	}
}

func main() {
	//创建具体的建造者
	builder := ConcreteBuilder{}
	//指定产品属性来创建产品
	people := builder.SetClothes("衣服1").SetHat("帽子1").SetPants("裤子1").SetSock("袜子1").SetShoes("鞋子1").Build()
	//打印产品
	people.ShowPeopleInfo()
}