package common

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 创建 redis 客户端
func RedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	fmt.Println("connect", pong)
	return client, nil
}
