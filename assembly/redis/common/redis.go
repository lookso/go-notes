package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// 创建 redis 客户端
func RedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		PoolSize:     40,              // Redis连接池大小
		MaxRetries:   3,               // 最大重试次数
		IdleTimeout:  5 * time.Second, // 空闲链接超时时间
		MinIdleConns: 5,               // 空闲连接数量
	})
	ctx := context.Background()
	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	fmt.Println("connect", pong)
	return client, nil
}
