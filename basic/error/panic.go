/*
@Time : 2019-11-04 21:29 
@Author : Tenlu
@File : panic
@Software: GoLand
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	type Node struct {
		Index int
		Value int
	}
	//var Nodes = make([]Node, 0)
	//Nodes = append(Nodes, Node{Index: 1, Value: 2})

	var Nodes = &[]Node{
		{Index: 1, Value: 1},
		{Index: 1, Value: 1},
		{Index: 2, Value: 2},
		{Index: 3, Value: 3},
		{Index: 5, Value: 5},
	}
	sum := 0
	res := make([]*int, 0)
	for _, v := range *Nodes {
		if v.Value > 1 {
			res = append(res, &v.Value)
		}
	}

	for _, v := range res {
		fmt.Println(*v) // 5 5 5 
		sum = +sum + *v
	}
	fmt.Println("sum", sum) // 15

	fmt.Println(catchErr() == nil)

	defer recover2()
	panic("some thing wrong")
}
func recover2() {
	recover1()
}

func recover1() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func catchErr() error {
	var err os.PathError
	return &err
}
