package main

import "fmt"

// 类的组合继承
type People struct {

}

type Teacher struct {
	People
}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

func main()  {
	var t=Teacher{}
	t.ShowA()
}