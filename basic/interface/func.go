package main

import (
	"fmt"
	"strconv"
)

type IOurs interface {
	She() string
}

type Our struct {
	Class string `json:"class"`
}

func (o *Our) She() string {
	return o.Class
}

func NewOurs(class string) *Our {
	return &Our{
		Class: class,
	}
}

type IMy interface {
	Human(o IOurs) (*Our, string)
	GetDay() int
	SetDays(day int) func() int // 生活了多少天
}
type My struct {
	Height int `json:"height"`
	Days   int `json:"days"`
	Ours   IOurs
}

func (m *My) Human(ours IOurs) (*Our, string) {
	return NewOurs("me"), ours.She() + ":" + strconv.Itoa(m.Days)
}

func (m *My) GetDay() int {
	return m.Days
}

func (m *My) SetDays(day int) func() int {
	f := func() int {
		m.Days = m.Days + day*m.Height
		return m.Days
	}
	f()
	return f
}

func NewMy() IMy {
	return &My{Height: 110, Days: 10, Ours: NewOurs("he")}
}

func Your(my IMy) {
	fmt.Println(my.GetDay())
	fmt.Println()
}

func main() {
	md := NewMy()
	md.SetDays(22)
	fmt.Println(md.GetDay())

	Your(md)
	fmt.Println(md.Human(NewOurs("she")))
	our, _ := md.Human(NewOurs("your"))
	fmt.Println(our.Class)

}
