package main

import "fmt"

// 中介模式
// https://juejin.cn/post/7003923907111551013

// https://www.cnblogs.com/ricklz/p/15631929.html

// 比如美国和伊拉克的安全问题，联合国安全理事对其进行协调。
type Country interface {
	SendMess(message string)
	GetMess(message string)
}

type USA struct {
	mediator *UnitedNationsSecurityCouncil
}

func (usa *USA) SendMess(message string) {
	usa.mediator.ForwardMessage(usa, message)
}

func (usa *USA) GetMess(message string) {
	fmt.Println("USA 获得对方的消息：", message)
}

type Irap struct {
	mediator *UnitedNationsSecurityCouncil
}

func (ir *Irap) SendMess(message string) {
	ir.mediator.ForwardMessage(ir, message)
}

func (ir *Irap) GetMess(message string) {
	fmt.Println("Irap 获得对方的消息：", message)
}

type Mediator interface {
	ForwardMessage(country Country, message string)
}

type UnitedNationsSecurityCouncil struct {
	USA
	Irap
}

func (uns *UnitedNationsSecurityCouncil) ForwardMessage(country Country, message string) {
	switch country.(type) {
	case *USA:
		uns.Irap.GetMess(message)
	case *Irap:
		uns.USA.GetMess(message)
	default:
		fmt.Println("国家不在联合国")
	}
}

func main()  {
	// 创建中介者，联合国
	mediator := &UnitedNationsSecurityCouncil{}

	usa := USA{mediator}
	mediator.USA = usa
	usa.SendMess("不准研制核武器，否则要发动战争了")

	irap := Irap{mediator}
	mediator.Irap = irap
	irap.SendMess("我们没有核武器")
}