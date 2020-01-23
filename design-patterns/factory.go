package main

import "fmt"

type Op interface {
	Connect() string // 类似php的interface Op{ public function Connect();} //接口内的方法必须实现
}
type TheMysql struct {
}
type ThePdo struct {
}
type TheFactory struct {
}

func (m *TheMysql) Connect() string {
	return "this is mysql drive connect"
}
func (p *ThePdo) Connect() string {
	return "this is pdo drive connect"
}
func (f *TheFactory) MyFactory(name string) Op {
	switch name {
	case "mysql":
		return new(TheMysql)
	case "pdo":
		return new(ThePdo)
	default:
		panic("name is error")
	}
	return nil
}
func main() {
	var f = new(TheFactory)
	p := f.MyFactory("mysql")
	fmt.Println(p.Connect())
}
