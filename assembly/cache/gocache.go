package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)
// 是一个运行在单机上的k/v缓存，相当于memcached，实现线程安全，可以带有过期时间访问清理
// https://blog.csdn.net/EDDYCJY/article/details/116725399
// http://www.iamlintao.com/7228.html
type My struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var my = My{
	Name: "jack",
	Age:  11,
}

func main() {
	// 创建一个cache对象，默认ttl 5分钟，每10分钟对过期数据进行一次清理
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set一个KV，key是"foo"，value是"bar"
	// TTL是默认值（上面创建对象的入参，也可以设置不同的值）5分钟
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set了一个没有TTL的KV，只有调用delete接口指定key时才会删除
	c.Set("baz", 42, cache.NoExpiration)

	// 从cache中获取key对应的value
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	// 如果想提高性能，存储指针类型的值
	c.Set("my_key", &my, cache.DefaultExpiration)
	if x, found := c.Get("my_key"); found {
		myKey := x.(*My)
		fmt.Println("my_key", myKey)
	}
}