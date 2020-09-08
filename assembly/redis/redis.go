package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("This is a program for go to use go_redis.")
	//connect
	client := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	defer client.Close()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

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
