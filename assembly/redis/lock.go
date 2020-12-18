package main

import (
	"fmt"
	"github.com/go-basic/uuid"
	"github.com/go-redis/redis"
	"sync"
	"time"
) //redis package

//connect redis
var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func init() {
	if err := client.Ping().Err(); err != nil {
		panic(err)
	}
}
func getUuid() string {
	uuid := uuid.New()
	return uuid
}

//lock
func lock(myfunc func()) {
	var lockKey = "my-lockr"
	//lock
	uuid := getUuid()
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
	value, _ := client.Get(lockKey).Result()
	if value == uuid { //compare value,if equal then del
		_, err := client.Del(lockKey).Result()
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
	time.Sleep(1 * time.Second)
	fmt.Printf("after incr is %d\n", counter)
}

//5 goroutine compete lock
var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			lock(incr)
		}()
	}
	wg.Wait()
	fmt.Printf("final counter is %d \n", counter)
}
