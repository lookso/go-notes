package main

import "fmt"

type MyUser struct {
	Name string `json:"name"`
}
type IUser interface {
	GetName() string
}

func (s *MyUser) GetName() string {
	return s.Name
}
func NewUser() IUser {
	return &MyUser{
		Name: "jack",
	}
}

func main() {
	u := NewUser()
	fmt.Println(u.GetName())
}