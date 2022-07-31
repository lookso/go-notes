package main

import "fmt"

type DbFactor interface {
	Connect() string
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
func (f *TheFactory) MyFactory(name string) DbFactor {
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
