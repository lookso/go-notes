package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

var stopCh = make(chan int)
var removeCh = make(chan int)

func TestOne() {
	var c = cron.New(cron.WithSeconds())
	entryID, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Printf("time in %v id:%+v,spec:%+v\n", time.Now().Unix(), 1, "*/5 * * * * *")
		stopCh <- 1
	})
	entryID2, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Printf("time in %v id:%+v,spec:%+v\n", time.Now().Unix(), 2, "*/2 * * * * *")
	})
	fmt.Println(entryID, err)
	fmt.Println(entryID2, err)
	c.Start()
	for {
		select {
		case <-stopCh:
			c.Stop()
		default:
		}
	}
}

func TestTwo() {
	var taskIDEntryIDMap = make(map[int]cron.EntryID)
	var c = cron.New(cron.WithSeconds())
	entryID, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Printf("time in %v id:%+v,spec:%+v\n", time.Now().Unix(), 1, "*/5 * * * * *")
		if time.Now().Unix() > 1668493121 {
			removeCh <- 1
		}
	})
	taskIDEntryIDMap[1] = entryID
	entryID2, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Printf("time in %v id:%+v,spec:%+v\n", time.Now().Unix(), 2, "*/5 * * * * *")
	})
	taskIDEntryIDMap[2] = entryID2

	fmt.Println(entryID, err)
	fmt.Println(entryID2, err)
	c.Start()
	for {
		fmt.Println(888)
		select {
		case taskID := <-removeCh:
			c.Remove(taskIDEntryIDMap[taskID])
		}
	}
	fmt.Println(6666)
}

var taskIDEntryIDMap = make(map[int]cron.EntryID)
var c = cron.New(cron.WithSeconds())

var taskStatusMap = make(map[int]*int)

func TestThree(spec string, taskID int) (cron.EntryID, error) {
	entryID, err := c.AddFunc(spec, func() {
		fmt.Printf("time in %v id:%+v,spec:%+v\n", time.Now().Unix(), taskID, spec)
		if _, ok := taskStatusMap[taskID]; ok {
			if *taskStatusMap[taskID] == 1 {
				removeCh <- taskID
			}
		}
	})
	fmt.Printf("entryID:%+v", entryID)
	return entryID, err
}
func main() {
	gin.SetMode(gin.DebugMode) // 设置在gin.New()前面才有效

	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/get/json", func(c *gin.Context) {

		id, ok := c.GetQuery("id")
		if !ok {
			return
		}
		if id != "" {
			var status *int
			num := 1
			status = &num
			taskStatusMap[cast.ToInt(id)] = status
		}
		c.JSON(http.StatusOK, id)
	})
	go GoTest()
	router.Run(":8000")
}

func GoTest() {

	tid := 1
	spec1 := "*/5 * * * * *"
	entryID, err := TestThree(spec1, tid)
	if err != nil {
		fmt.Println(err)
	} else {
		taskIDEntryIDMap[tid] = entryID
	}
	var status *int
	num := 0
	status = &num
	taskStatusMap[tid] = status

	fmt.Println("1---", entryID, err)
	tid = 2
	spec2 := "*/5 * * * * *"
	entryID2, err := TestThree(spec2, tid)
	if err != nil {
		fmt.Println(err)
		fmt.Println("malegebi")
	} else {
		taskIDEntryIDMap[tid] = entryID2
	}
	var status1 *int
	num1 := 0
	status1 = &num1
	taskStatusMap[tid] = status1

	fmt.Println("2---", entryID2, err)
	c.Start()

	for {
		select {
		case taskID := <-removeCh:
			c.Remove(taskIDEntryIDMap[taskID])
		//fmt.Println(taskID)
		default:
			time.Sleep(time.Second)
			fmt.Println(88888)

		}
	}
}
