package main

import "fmt"

type Action interface {
	Run() string
	GetDay() int
	SetDays(day int) func() int // 生活了多少天
}

type MyDo struct {
	Height int `json:"height"`
	Days   int `json:"days"`
}

func (*MyDo) Run() string {
	panic("implement me")
}

func (m *MyDo) GetDay() int {
	return m.Days
}

func (m *MyDo) SetDays(day int) func() int {
	f := func() int {
		m.Days = day * m.Height
		return m.Days
	}
	f()
	return f
}

func NewMyDo() Action {
	return &MyDo{Height: 110}
}

func main() {
	md := NewMyDo()
	md.SetDays(22)
	fmt.Println(md.GetDay())

}
