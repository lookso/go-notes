package main

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"go-notes/assembly/redis/common"
	"time"
)

// 100s内最多允许50个请求
func main() {
	rate()
}

func rate() {
	client, _ := common.RedisClient()
	key := fmt.Sprintf("user_id_%v", 10086)
	ctx := context.Background()
	len, _ := client.LLen(ctx, key).Result()
	if len < 5 {
		err := client.LPush(ctx, key, time.Now().Unix()).Err()
		if err != nil {
			fmt.Println("err", err)
		}
	} else {
		t, _ := client.LIndex(ctx, key, -1).Result()
		if time.Now().Unix()-cast.ToInt64(t) < 100 {
			fmt.Println("请求过多，请稍后再试")
		} else {
			err := client.LPush(ctx, key, time.Now().Unix()).Err()
			if err != nil {
				fmt.Println("err", err)
			}
			// ltrim 裁剪 让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除
			if err := client.LTrim(ctx, key, 0, 50).Err(); err != nil {
				fmt.Println("err", err)
			}
		}
	}
}
