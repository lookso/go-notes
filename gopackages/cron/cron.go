package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

// 声明一个全局的redisDb变量
var redisDb *redis.Client

// 根据redis配置初始化一个客户端
func initClient() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

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

var taskStatusMap = make(map[int]*int) // 1 暂停

func TestThree(spec string, taskID int) (cron.EntryID, error) {
	entryID, err := c.AddFunc(spec, func() {
		fmt.Printf("time in %v id:%+v,spec:%+v\n", time.Now().Format("2006-01-02 15:04:05"), taskID, spec)
		if _, ok := taskStatusMap[taskID]; ok {
			if *taskStatusMap[taskID] == 1 {
				removeCh <- taskID
			}
		}
	})
	fmt.Printf("cron entryID:%+v-", entryID)
	return entryID, err
}
func main() {
	if err := initClient(); err != nil {
		panic(err)
	}

	gin.SetMode(gin.DebugMode) // 设置在gin.New()前面才有效

	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/get/json", func(c *gin.Context) {

		//id, ok := c.GetQuery("id")
		taskSpec, _ := c.GetQuery("task_spec")
		taskStatus, _ := c.GetQuery("task_status")

		id := 1
		if cast.ToInt(taskStatus) > 0 {
			var status *int
			setStatus := cast.ToInt(taskStatus)
			status = &setStatus
			taskStatusMap[cast.ToInt(id)] = status
		}


		if taskSpec != "" {
			redisDb.Set("1", taskSpec, 0)
		}
		fmt.Printf("http cron change time:%+v,taskSpec:%+v\n", time.Now().Format("2006-01-02 15:04:05"), taskSpec)
		c.JSON(http.StatusOK, id)
	})
	go goRunCron()
	router.Run(":8000")
}

func goRunCron() {

	tid := 1
	spec := "0 * * * * *"
	fmt.Println("taskStatusMap[tid]", taskStatusMap[tid])

	var newStatus *int
	newSetStatus := 0
	newStatus = &newSetStatus
	taskStatusMap[tid] = newStatus

	newSpec, err := redisDb.Get("1").Result()
	fmt.Printf("newSpec-111:%+v", newSpec)
	if newSpec != "" {
		spec = newSpec
	}

	entryID, err := TestThree(spec, tid)
	if err != nil {
		fmt.Printf("run err:%+v", err)
	} else {
		taskIDEntryIDMap[tid] = entryID
	}
	fmt.Printf("after spec:%+v,time:%+v,entryID:%+v,err:%+v\n", spec, time.Now().Format("2006-01-02 15:04:05"), entryID, err)

	spec = "* * * * * *"
	c.Start()

	for {
		select {
		case taskID := <-removeCh:
			fmt.Println("remove taskid:%+v", taskID)
			c.Remove(taskIDEntryIDMap[taskID])
		default:
			//time.Sleep(time.Second)
			//fmt.Println(88888)
		}
	}
}
