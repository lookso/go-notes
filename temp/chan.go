package main

import (
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
)

type Task struct {
	index int
	nums  []int
	sum   int
	wg    *sync.WaitGroup
}

func (t *Task) Do() {
	for _, num := range t.nums {
		t.sum += num
	}
	//tt := 1 / 0
	panic("cuowu")

	t.wg.Done()
}

func panicHandler(err interface{}) {
	defer func() {
		err = recover()
		fmt.Println(err)
	}()
}

func main() {
	env := "gray"
	if env == "gray" {
		env = "online"
	}
	fmt.Println("env", env)

	c := context.Background()
	c.Value("123")

	fmt.Println(fmt.Sprintf("a%s", "nb"))

	var wg sync.WaitGroup
	wg.Add(1)
	p, _ := ants.NewPoolWithFunc(30, taskFunc, ants.WithPanicHandler(panicHandler))
	var t = &Task{
		index: 1,
		nums:  []int{1, 2, 4},
		sum:   10,
		wg:    &wg,
	}
	p.Invoke(t)
	wg.Wait()
	defer p.Release()

}
func taskFunc(data interface{}) {
	task := data.(*Task)
	task.Do()
	fmt.Printf("task:%d sum:%d\n", task.index, task.sum)
}
