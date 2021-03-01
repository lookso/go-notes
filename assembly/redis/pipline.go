package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-notes/assembly/redis/common"
)

func main() {
	client,_ := common.RedisClient()
	client.HSet("1371648200","sex",100)
	pipe := client.Pipeline()
	pipe.HGetAll("1371648200")
	pipe.HGetAll("1371648200")
	cmders, err := pipe.Exec()
	if err != nil {
		fmt.Println("err", err)
	}
	for _, cmder := range cmders {
		cmd := cmder.(*redis.StringStringMapCmd)
		strMap, err := cmd.Result()
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("strMap", strMap)
	}
}