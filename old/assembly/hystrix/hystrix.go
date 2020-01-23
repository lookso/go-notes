package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"time"
)

func main() {

	hystrix.ConfigureCommand("get_baidu", hystrix.CommandConfig{
		Timeout:                5,  //超时时间设置  单位毫秒
		MaxConcurrentRequests:  100,  //最大请求数
		RequestVolumeThreshold: 50000,   //过多长时间，熔断器再次检测是否开启。单位毫秒
		SleepWindow:            2,    //请求阈值,熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
		ErrorPercentThreshold:  3, //错误率
	})
	var j = 0
	for i := 0; i < 100; i++ {
		TestHystix()
		fmt.Println("nums:", j)
		j++
		time.Sleep(1 * time.Second)
	}
	time.Sleep(2 * time.Second) // 调用Go方法就是起了一个goroutine，这里要sleep一下，不然看不到效果
}

func TestHystix() {

	// 根据自身业务需求封装到http client调用处
	hystrix.Go("get_baidu", func() error {
		// 调用关联服务
		res, err := http.Get("https://www.jianshu.com/p/135526f306fb")
		if err != nil {
			fmt.Println("get error")
			return err
		}
		fmt.Println("请求成功：", res.Status)
		return nil
	},
		// 失败重试，降级等具体操作
		func(err error) error {
			fmt.Println("get an error, handle it")
			return nil
		})
}
