package main

import (
	"context"
	"fmt"
	"go-notes/assembly/redis/common"
)

func main() {
	ctx:=context.Background()
	client, _ := common.RedisClient()
	// set key1 first
	// set key2 second
	res, err := client.Eval(ctx,"redis.call('MSET', KEYS[1],ARGV[1],KEYS[2],ARGV[2]) return redis.call('MGET', KEYS[1],KEYS[2])", []string{"key1", "key2"}, "first", "second").Result()
	fmt.Println(res, err)
}
