/*
@Time : 2019-10-22 17:36
@Author : Tenlu
@File : download
@Software: GoLand
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Run() (int, error)
}

type DownLoadTask struct {
	File string
}

func (d *DownLoadTask) Run() (int, error) {
	return time.Now().Nanosecond(), nil
}

func Tasks() []Task {
	task1 := &DownLoadTask{File: "task1.txt"}
	task2 := &DownLoadTask{File: "task2.txt"}
	task3 := &DownLoadTask{File: "task3.txt"}
	workList := []Task{task1, task2, task3}
	return workList
}

func main() {
	
	tasks := Tasks()
	taskLen := len(tasks)
	var wg sync.WaitGroup
	wg.Add(taskLen)

	for k, task := range tasks {
		go func(num int) {
			if rtime, err := task.Run(); err == nil {
				fmt.Printf("runtime:%d\n", rtime)
			}
			wg.Done()
		}(k)
	}
	wg.Wait()
}
