package main

import "fmt"

func init() {
	initA()
	initB()
}
func initA() {
	fmt.Println(123)
}
func initB() {
	fmt.Println(456)
}

type Cp struct {
	Name string `json:"name"`
	Sex  int    `json:"sex"`
}

func hd() {
	var cp *Cp
	tag := 1
	for i := 0; i < 3; i++ {
		if tag == 1 {
			cp = &Cp{
				Name: "jack",
				Sex:  1,
			}
		}
		tag = 2
		fmt.Println(cp.Name)
	}
}
func main() {
	hd()
}
