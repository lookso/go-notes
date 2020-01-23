/*
@Time : 2019-07-24 20:52
@Author : Tenlu
@File : composite
@Software: GoLand
*/
package main

import "fmt"

//意图：将对象组合成树形结构以表示"部分-整体"的层次结构。组合模式使得用户对单个对象和组合对象的使用具有一致性。
//主要解决：它在我们树型结构的问题中，模糊了简单元素和复杂元素的概念，客户程序可以像处理简单元素一样来处理复杂元素，
//从而使得客户程序与复杂元素的内部结构解耦。

// 菜单描述
type MenuDesc struct {
	name        string
	description string
}

func (m *MenuDesc) Name() string {
	return m.name
}
func (m *MenuDesc) Description() string {
	return m.description
}


// 菜单组
type MenuItem struct {
	MenuDesc
	price float32  //套菜组价格
}

func (m *MenuItem) Price() float32 {
	return m.price
}

func (m *MenuItem) Print() {
	fmt.Printf("  %s, ￥%.2f\n", m.name, m.price)
	fmt.Printf("  -- %s\n", m.description)
}
func NewMenuItem(name, description string, price float32) *MenuItem {
	return &MenuItem{
		MenuDesc: MenuDesc{
			name:        name,
			description: description,
		},
		price: price,
	}
}



type MenuComponent interface {
	Price() float32
	Print()
}

type MenuGroup struct {
	children []MenuComponent
}

type Menu struct {
	MenuDesc
	MenuGroup
}
// 添加菜单组合详情
func (m *Menu) Add(c MenuComponent) {
	m.children = append(m.children, c)
}

func (m *Menu) Remove(idx int) {
	m.children = append(m.children[:idx], m.children[idx+1:]...)
}

func (m *Menu) Child(idx int) MenuComponent {
	return m.children[idx]
}

func (m *Menu) Price() (price float32) {
	for _, v := range m.children {
		price += v.Price()
	}
	return
}
func (m *Menu) Print() {
	fmt.Printf("%s, %s, ￥%.2f\n", m.name, m.description, m.Price())
	fmt.Println("------------------------")
	for _, v := range m.children {
		v.Print()
	}
}
func NewMenu(name, description string) *Menu {
	return &Menu{
		MenuDesc: MenuDesc{
			name:        name,
			description: description,
		},
	}
}
func main() {
	menu1 := NewMenu("培根鸡腿燕麦堡套餐", "供应时间：09:15--22:44")
	menu1.Add(NewMenuItem("主食", "培根鸡腿燕麦堡1个", 11.5))
	menu1.Add(NewMenuItem("小吃", "玉米沙拉1份", 5.0))
	menu1.Add(NewMenuItem("饮料", "九珍果汁饮料1杯", 6.5))

	menu2 := NewMenu("奥尔良烤鸡腿饭套餐", "供应时间：09:15--22:44")
	menu2.Add(NewMenuItem("主食", "新奥尔良烤鸡腿饭1份", 15.0))
	menu2.Add(NewMenuItem("小吃", "新奥尔良烤翅2块", 11.0))
	menu2.Add(NewMenuItem("饮料", "芙蓉荟蔬汤1份", 4.5))

	all := NewMenu("超值午餐", "周一至周五有售")
	all.Add(menu1)
	all.Add(menu2)

	all.Print()
}

//超值午餐, 周一至周五有售, ￥53.50
//------------------------
//培根鸡腿燕麦堡套餐, 供应时间：09:15--22:44, ￥23.00
//------------------------
//主食, ￥11.50
//-- 培根鸡腿燕麦堡1个
//小吃, ￥5.00
//-- 玉米沙拉1份
//饮料, ￥6.50
//-- 九珍果汁饮料1杯
//奥尔良烤鸡腿饭套餐, 供应时间：09:15--22:44, ￥30.50
//------------------------
//主食, ￥15.00
//-- 新奥尔良烤鸡腿饭1份
//小吃, ￥11.00
//-- 新奥尔良烤翅2块
//饮料, ￥4.50
//-- 芙蓉荟蔬汤1份
