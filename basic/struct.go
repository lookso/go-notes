package main

import "fmt"

type user struct {
	age int
}
type family struct {
	user
}

func main() {
	var u user
	u.addAge(10)
	u.lessAge(5)
	fmt.Println(u.getAge())
}

func (u *user) addAge(n int) {
	u.age += n
}
func (u *user) lessAge(n int) {
	u.age -= n
}
func (u *user) getAge() string {
	return fmt.Sprintf("%v", u.age)
}
func (u *user) getName() string {
	return "u"
}
func (f *family) getName() string {
	return f.user.getName()
}
func (u *user) Name() {
}
