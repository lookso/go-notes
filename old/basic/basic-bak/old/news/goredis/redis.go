/*
@Time : 2019/4/6 7:14 PM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
	"github.com/pkg/errors"
)

func main()  {
	fmt.Println("This is a program for go to use go_redis.")
	client:=createClient()
	println("-------String---------\n")
	redisString(client)
	println("-------Hash---------\n")
	redisHash(client)
	println("-------List---------\n")
	redisList(client)
	println("-------set---------\n")
	redisSet(client)
	println("-------zset---------\n")
	redisZset(client)
}

func redisString(client *redis.Client)  {
	name,err:=client.Get("company").Result()
	if err==redis.Nil{
		fmt.Println("key not exists")
		err:=client.MSet("company","alibaba","address","hangzhou").Err()
		if err!=nil{
			fmt.Println(err)
		}
	}else{
		fmt.Println(err)
	}
	num,_:=client.Incr("count").Result()
	fmt.Printf("num is %d\n",num)
	addLen,_:=client.StrLen("address").Result()
	fmt.Printf("`address`'s strlen is %d\n",addLen)
	fmt.Println(name)

	if 0 < time.Second || 0%time.Second != 0{
		fmt.Printf("the Second:%s\n",time.Second)
		fmt.Printf("the Millisecond:%s\n ",20*time.Millisecond)
	}
	phoneErr:=client.Set("phoneno","15010908928",time.Second*20).Err()
	if phoneErr==redis.Nil{
		fmt.Printf("phoneno empty\n")
	}
	count:=client.Incr("count").Val()
	fmt.Printf("%d\n",count)

	client.SetBit("zannum",10086,1).Err()
	client.SetBit("zannum",96103,1).Err()
	zannum:=client.GetBit("zannum",10086).Val()
	fmt.Printf("count num:%d\n",client.BitCount("zannum",&redis.BitCount{0,-1}).Val())
	fmt.Printf("zannum is %d\n",zannum)
}
func redisHash(client *redis.Client)  {
	name,err:=client.HGet("uid_10086","name").Result()
	if err==redis.Nil{
		err:=client.HMSet("uid_10086",map[string]interface{}{"name":"jacks","age":1}).Err()
		if err!=nil{
			errors.New("hset error")
		}
	}else {
		fmt.Printf("name is %s\n", name)
		uidInfo,_:=client.HGetAll("uid_10086").Result()
		for k,v:=range uidInfo {
			fmt.Printf("%s:%s\n",k,v)
		}
		uidlen:=client.HLen("uid_10086").Val()
		fmt.Printf("len:%d\n",uidlen)
	}
}
func redisList(client *redis.Client)  {
	isexists:=client.Exists("language").Val()
	if isexists==0 {
		err := client.LPush("language", "php", "golang", "python").Err()
		if err != nil {
			fmt.Println(err)
		}
	}
	language,_:=client.LRange("language",0,-1).Result()
	for _,l:=range language  {
		fmt.Printf("i like language is %s\n",l)
	}
	fmt.Println(client.LLen("language").Val())
	//for {
	//	pop:=client.BRPop(time.Second*20,"language").Val()
	//	fmt.Printf("pop ele is %s->%s\n",pop[0],pop[1])
	//}
}
func redisSet(client *redis.Client)  {
	err:=client.SAdd("yourlike","football","baskball","sleep").Err()
	if err!=nil{
		fmt.Println(err)
	}
	myLikeErr:=client.SAdd("mylike","coding","baskball","run").Err()
	if myLikeErr!=nil{
		fmt.Println(myLikeErr)
	}
	inter,_:=client.SInter("yourlike","mylike").Result()
	for _,like:=range inter {
		fmt.Printf("our inter like is %s\n",like)
	}
}

func redisZset(client *redis.Client)  {
	err:=client.ZAdd("height",redis.Z{Score:8848,Member:"zhufeng"}).Err()
	if err!=nil{
		fmt.Println(err)
	}
	zhufeng:=client.ZScore("height","zhufeng").Val()
	fmt.Printf("zhumulangmafeng height:%dm\n",int(zhufeng))
}

// 创建 redis 客户端
func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}