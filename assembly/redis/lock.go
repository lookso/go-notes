package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"log"
	"strconv"
	"sync"
	"time"
) //redis package

//connect redis
var client = redis.NewClient(&redis.Options{
	Addr:     ":6379",
	Password: "",
	DB:       0,
})

func init() {
	if err := client.Ping().Err(); err != nil {
		panic(err)
	}
}

func getUuid() string {
	return uuid.New().String()
}

//lock
func lock(myfunc func()) {
	uuid := getUuid()
	var lockKey = "mylockr"
	//lock
	lockSuccess, err := client.SetNX(lockKey, uuid, time.Second*5).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail")
		return
	} else {
		fmt.Println("get lock")
	}
	//run func
	myfunc()
	//unlock
	// get和del不是原子性,考虑放到lua脚本
	value, _ := client.Get(lockKey).Result()
	if value == uuid {
		_, err = client.Del(lockKey).Result()

		if err != nil {
			fmt.Println("unlock fail")
		} else {
			fmt.Println("unlock")
		}
	}
}

//do action
var counter int64

func incr() {
	counter++
	time.Sleep(6 * time.Second)
	fmt.Printf("after incr is %d\n", counter)
}

//5 goroutine compete lock
var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			lock(incr)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("final counter is %d \n", counter)
	lk()
}

func lk(){
	fmt.Println("This is a program for go to use go_redis.")
	//connect
	client,_:=RedisClient()
	defer client.Close()

	var rw sync.RWMutex
	var group sync.WaitGroup
	var j = 0
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func(i int) {
			defer group.Done()
			for {
				j++
				err := client.Do("set", "groupId:"+strconv.Itoa(i)+strconv.Itoa(j), "name:"+strconv.Itoa(i)+strconv.Itoa(j), "ex", "50", "nx").Err()
				if err != nil {
					log.Fatal("err", err)
				}
				name, _ := client.Get("groupId:" + strconv.Itoa(i) + strconv.Itoa(j)).Result()
				if name != "" {
					fmt.Println("string:", name)
				}
				rw.RLock()
			}
		}(i)
	}
	group.Wait()

	select {}
}