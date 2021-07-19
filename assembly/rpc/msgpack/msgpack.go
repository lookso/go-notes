package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
	"reflect"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
	fmt.Println("原数据-", countryCapitalMap)
	jsBt, err := json.Marshal(&countryCapitalMap)
	fmt.Println("jsStr", string(jsBt))
	//in := map[string]interface{}{"foo": uint32(123456789), "hello": "world"}
	in := countryCapitalMap
	res, err := msgpack.Marshal(in)
	if err != nil {
		fmt.Printf("序列化失败")
	}
	fmt.Println("reflect.TypeOf", reflect.TypeOf(res))
	fmt.Println("序列化数据--", res)

	//连接redis
	initClient()
	//存入redis数据类型[]type可以存入
	bool := rdb.Set("val", res, 0).Err()
	if bool != nil {
		fmt.Printf("set val failed, err:%v\n", err)
		return
	}
	//返回类型可变
	val, err := rdb.Get("val").Bytes()
	if err != nil {
		fmt.Printf("get val failed, err:%v\n", err)
		return
	}
	fmt.Println("redis取出数据--", val)
	var out map[string]string
	bool = msgpack.Unmarshal(val, &out)
	if bool != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println("反序列化数据--", out)
}
