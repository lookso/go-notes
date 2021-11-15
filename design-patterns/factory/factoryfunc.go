// 工厂方法模式
// 通过该种方式，可以自由的实现对产品的扩展。添加新的手机时，实现实现Phone接口，在实现一个创建该产品的方法，需要实现Factory接口
// https://cloud.tencent.com/developer/article/1365764
package main

import "fmt"

type Phone interface {
	ShowBrand()
}

type Factory interface {
	CreatePhone() Phone
}

// iphone
type IPhone struct {
}

func (p *IPhone) ShowBrand() {
	fmt.Println("我是苹果手机")
}

// 华为
type HPhone struct {
}

func (p *HPhone) ShowBrand() {
	fmt.Println("我是华为手机")
}

// 小米
type XPhone struct {
}

func (p *XPhone) ShowBrand() {
	fmt.Println("我是小米手机")
}

// 华为工厂
type HFactory struct {
}

func (F *HFactory) CreatePhone() Phone {
	return &HPhone{}
}

// 苹果工厂
type IFactory struct {
}

func (F *IFactory) CreatePhone() Phone {
	return &IPhone{}
}

// 小米工厂
type XFactory struct {
}

func (F *XFactory) CreatePhone() Phone {
	return &XPhone{}
}

func main() {
	var phone Phone

	// 苹果工厂实例
	appleFactory := &IFactory{}
	phone = appleFactory.CreatePhone()
	phone.ShowBrand()

	// 小米手机工厂
	xmFactory := &XFactory{}
	phone = xmFactory.CreatePhone()
	phone.ShowBrand()

	// 华为手机
	huaweiFactory := &HFactory{}
	phone = huaweiFactory.CreatePhone()
	phone.ShowBrand()
}