package test

import (
	"context"
	"fmt"
	"go-notes/assembly/redis/common"
	"log"
	"strconv"
	"sync"
	"testing"
)

func TestGoRedis(t *testing.T) {
	ctx := context.Background()
	fmt.Println("This is a program for go to use go_redis.")
	//connect
	client, _ := common.RedisClient()
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
				err := client.Do(ctx, "set", "groupId:"+strconv.Itoa(i)+strconv.Itoa(j), "name:"+strconv.Itoa(i)+strconv.Itoa(j), "ex", "50", "nx").Err()
				if err != nil {
					log.Fatal("err", err)
				}
				name, _ := client.Get(ctx, "groupId:"+strconv.Itoa(i)+strconv.Itoa(j)).Result()
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
