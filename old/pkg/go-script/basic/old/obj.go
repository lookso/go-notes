package main

import "fmt"

type People interface {
	run()
}

// 多态
type Man interface {
	run()
}
type Women interface {
	run()
}
type father struct {
	name    string
	job     string
	country string
	sex     int
}
type son struct {
	*father        // 继承
	name    string // 重载
	job     string
	height  float32
}

func (s *son) run() string {
	return s.job
}
func main() {
	defer fly()
	s := &son{&father{"toms", "fruit-growers", "china", 1}, "jack", "ceo", 173.21}
	fmt.Printf("son'name is: %s,son's job:%s,sex is: %d\n", s.name, s.run(), s.sex)
	fmt.Printf("father's name is:%s,father's job is:%s,sex is:%d,father's son is %s\n", s.father.name, s.father.job, s.father.sex, s.name)
}

func fly() { // 方法名首字母小写包内部调用类似private,首字母大写则外部调用类似public
	fmt.Printf("i can fly\n")
}
func ToFly() { // 包外调用
	fmt.Printf("you can fly\n")
}
