package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
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
}
