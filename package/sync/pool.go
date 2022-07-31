package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
}

var pool *sync.Pool

func main() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person")
			return new(Person)
		},
	}
	p := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p)

	p.Name="jack"

	fmt.Printf("设置 p.Name = %s\n", p.Name)

	pool.Put(p)

	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", pool.Get().(*Person))
	fmt.Println("Pool 没有对象了，调用 Get: ", pool.Get().(*Person))
}
