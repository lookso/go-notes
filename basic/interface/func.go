package main

import "fmt"

type Action interface {
	UserInfo() string
	GetDay() int
	SetDays(day int) func() int // 生活了多少天
}

type MySelf struct {
	Height int `json:"height"`
	Days   int `json:"days"`
}

func (m *MySelf) UserInfo() string {
	panic("implement me")
}

func (m *MySelf) GetDay() int {
	return m.Days
}

func (m *MySelf) SetDays(day int) func() int {
	f := func() int {
		m.Days = m.Days+day * m.Height
		return m.Days
	}
	f()
	return f
}

func NewMyDo() Action {
	return &MySelf{Height: 110,Days:10}
}

func main() {
	md := NewMyDo()
	md.SetDays(22)
	fmt.Println(md.GetDay())

}
