/*
@Time : 2019/4/6 7:14 PM
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"go-notes/assembly/redis/common"
	"time"
)

func main() {
	ctx := context.Background()
	fmt.Println("This is a program for go to use go_redis.")
	client, _ := common.RedisClient()
	println("-------String---------\n")
	redisString(ctx, client)
	println("-------Hash---------\n")
	redisHash(ctx, client)
	println("-------List---------\n")
	redisList(ctx, client)
	println("-------set---------\n")
	redisSet(ctx, client)
	println("-------zset---------\n")
	redisZset(ctx, client)
}

func redisString(ctx context.Context, client *redis.Client) {

	name, err := client.Get(ctx, "company").Result()
	if err == redis.Nil {
		fmt.Println("key not exists")
		err := client.MSet(ctx, "company", "alibaba", "address", "hangzhou").Err()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
	num, _ := client.Incr(ctx, "count").Result()
	fmt.Printf("num is %d\n", num)
	addLen, _ := client.StrLen(ctx, "address").Result()
	fmt.Printf("`address`'s strlen is %d\n", addLen)
	fmt.Println(name)

	if 0 < time.Second || 0%time.Second != 0 {
		fmt.Printf("the Second:%s\n", time.Second)
		fmt.Printf("the Millisecond:%s\n ", 20*time.Millisecond)
	}
	phoneErr := client.Set(ctx, "phoneno", "15010908928", time.Second*20).Err()
	if phoneErr == redis.Nil {
		fmt.Printf("phoneno empty\n")
	}
	count := client.Incr(ctx, "count").Val()
	fmt.Printf("%d\n", count)

	client.SetBit(ctx, "zannum", 10086, 1).Err()
	client.SetBit(ctx, "zannum", 96103, 1).Err()
	zannum := client.GetBit(ctx, "zannum", 10086).Val()
	fmt.Printf("count num:%d\n", client.BitCount(ctx, "zannum", &redis.BitCount{0, -1}).Val())
	fmt.Printf("zannum is %d\n", zannum)
}
func redisHash(ctx context.Context, client *redis.Client) {
	name, err := client.HGet(ctx, "uid_10086", "name").Result()
	if err == redis.Nil {
		err := client.HMSet(ctx, "uid_10086", map[string]interface{}{"name": "jacks", "age": 1}).Err()
		if err != nil {
			errors.New("hset error")
		}
	} else {
		fmt.Printf("name is %s\n", name)
		uidInfo, _ := client.HGetAll(ctx, "uid_10086").Result()
		for k, v := range uidInfo {
			fmt.Printf("%s:%s\n", k, v)
		}
		uidlen := client.HLen(ctx, "uid_10086").Val()
		fmt.Printf("len:%d\n", uidlen)
	}
}
func redisList(ctx context.Context, client *redis.Client) {
	isexists := client.Exists(ctx, "language").Val()
	if isexists == 0 {
		err := client.LPush(ctx, "language", "php", "golang", "python").Err()
		if err != nil {
			fmt.Println(err)
		}
	}
	language, _ := client.LRange(ctx, "language", 0, -1).Result()
	for _, l := range language {
		fmt.Printf("i like language is %s\n", l)
	}
	fmt.Println(client.LLen(ctx, "language").Val())
	//for {
	//	pop:=client.BRPop(time.Second*20,"language").Val()
	//	fmt.Printf("pop ele is %s->%s\n",pop[0],pop[1])
	//}
}
func redisSet(ctx context.Context, client *redis.Client) {
	err := client.SAdd(ctx, "yourlike", "football", "baskball", "sleep").Err()
	if err != nil {
		fmt.Println(err)
	}
	myLikeErr := client.SAdd(ctx, "mylike", "coding", "baskball", "run").Err()
	if myLikeErr != nil {
		fmt.Println(myLikeErr)
	}
	inter, _ := client.SInter(ctx, "yourlike", "mylike").Result()
	for _, like := range inter {
		fmt.Printf("our inter like is %s\n", like)
	}
}

func redisZset(ctx context.Context, client *redis.Client) {
	err := client.ZAdd(ctx, "height", &redis.Z{Score: 8848, Member: "zhufeng"}).Err()
	if err != nil {
		fmt.Println(err)
	}
	zhufeng := client.ZScore(ctx, "height", "zhufeng").Val()
	fmt.Printf("zhumulangmafeng height:%dm\n", int(zhufeng))
}
