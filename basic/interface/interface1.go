package main

import "fmt"

type IDao interface {
	GetAddr() string
	SetAddr(string) // 定义的接口方法需要全部实现
}
type Base struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}
type Mysql struct {
	Base
}

func (m *Mysql) GetAddr() string {
	return m.Name
}
func (m *Mysql) SetAddr(name string) {
	m.Name = name
}

func newMysql() *Mysql {
	var m Mysql
	m.Name = TypeMysql
	return &m
}

// =========================

type Redis struct {
	Base
}

func (m *Redis) GetAddr() string {
	return m.Name
}
func (m *Redis) SetAddr(name string) {
	m.Name = name
}

func newRedis() *Redis {
	var r Redis
	r.Name = TypeMysql
	return &r
}

const (
	TypeMysql = "mysql"
	TypeRedis = "redis"
)

func build(db string) IDao {
	switch db {
	case TypeMysql:
		return newMysql()
	case TypeRedis:
		return newRedis()
	default:
		panic("gcache: Unknown type ")
	}

}
func main() {
	var dbs = []string{"mysql", "redis"}
	for _, v := range dbs {
		b := build(v)
		b.SetAddr(v + ":127.0.0.1")
		fmt.Println("addr", b.GetAddr())
	}
}
