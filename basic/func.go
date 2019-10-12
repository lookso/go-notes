package main

import (
	"fmt"
	"time"
)
// 不仅仅结构体可以作为方法的声明,type指定的都可以,方法是包含了接收者的函数
type User map[string]string
type Handler func (name string) int

func (h Handler) add(name string) int {
	return h(name) + 10
}

func (u User) SetName(key string) error {
	u[key] = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
func (u User) GetName(key string) string {
	if v, ok := u[key]; ok {
		return v
	}
	return ""
}
func main() {
	var user = make(User)
	key := "nowTime"
	if err := user.SetName(key); err == nil {
		fmt.Println(user.GetName(key))
	}
}
