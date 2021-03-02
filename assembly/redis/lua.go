package main

import (
	"fmt"
	"go-notes/assembly/redis/common"
)

func main() {
	client, _ := common.RedisClient()
	// set key1 first
	// set key2 second
	res, err := client.Eval("return redis.call('MSET', KEYS[1],ARGV[1],KEYS[2],ARGV[2])", []string{"key1", "key2"}, "first", "second").Result()
	fmt.Println(res, err)
}
