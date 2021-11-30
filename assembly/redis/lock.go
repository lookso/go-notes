package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go-notes/assembly/redis/common"
	"time"
) //redis package

func getUuid() string {
	return uuid.New().String()
}

var client *redis.Client

func init() {
	client, _ = common.RedisClient()
}

//lock
func lock(ctx context.Context, myfunc func()) {
	uuid := getUuid()
	var lockKey = "mylockr"
	//lock
	lockSuccess, err := client.SetNX(ctx, lockKey, uuid, time.Second*5).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail")
		return
	} else {
		fmt.Println("get lock")
	}
	//run func
	myfunc()

	//unlock
	// get和del不是原子性,考虑放到lua脚本,使用如下脚本
	/*
	value, _ := client.Get(ctx, lockKey).Result()
	if value == uuid {
		_, err = client.Del(ctx, lockKey).Result()
		if err != nil {
			fmt.Println("unlock fail")
		} else {
			fmt.Println("unlock")
		}
	}
	*/
	var luaScript = redis.NewScript(`
        if redis.call("get", KEYS[1]) == ARGV[1]
            then
                return redis.call("del", KEYS[1])
            else
                return 0
        end
    `)
	rs, _ := luaScript.Run(context.Background(),client, []string{lockKey}, uuid).Result()
	if rs == 0 {
		fmt.Println("unlock fail")
	} else {
		fmt.Println("unlock")
	}
}
// https://mp.weixin.qq.com/s/Uo507elzCXzI0eI7mDHpGg

//do action
var counter int64

func incr() {
	time.Sleep(10 * time.Second)
	counter++
	fmt.Printf("after incr is %d\n", counter)
}

//5 goroutine compete lock
//var wg sync.WaitGroup

func main() {
	ctx := context.Background()
	for i := 0; i < 50; i++ {
		//wg.Add(1)
		//go func() {
			lock(ctx, incr)
		//	defer wg.Done()
	//	}()
	}
	//wg.Wait()
	fmt.Printf("final counter is %d \n", counter)
}
//https://hub.fastgit.org/zieckey/etcdsync/blob/master/mutex.go


