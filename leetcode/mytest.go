package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
)

type Params = map[string]interface{}

type Show struct {
	P Params
}

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("j: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("teoms")
	return

	s := new(Show)
	Params := make(map[string]interface{})
	Params["RMB"] = 100
	s.P = Params
	T1()
	T2()

}

type People struct {
	Name string `json:"name"`
}

func T1() {
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func T2() {
	var a int
	defer func() {
		fmt.Println("a:",a)
	}()
	a = 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	// defer 在定义的时候会计算好调用函数的参数，所以会优先输出10、20 两个参数。然后根据定义的顺序倒序执行。
}
//10 1 2 3
//20 0 2 2
//2 0 2 2
//1 1 3 4
