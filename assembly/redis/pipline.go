package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-notes/assembly/redis/common"
)

func main() {
	ctx := context.Background()
	client, _ := common.RedisClient()
	client.HSet(ctx, "1371648200", "sex", 100)
	pipe := client.Pipeline()
	pipe.HGetAll(ctx, "1371648200")
	pipe.HGetAll(ctx, "1371648200")
	cmders, err := pipe.Exec(ctx)
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
