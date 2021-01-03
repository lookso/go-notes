/*
@Time : 2019/5/22 11:02 AM 
@Author : Tenlu
@File : ok
@Software: GoLand
*/
package main

import "fmt"

func main()  {
	/**
	1.检测通道是否关闭
	2.检测map指定的key是否存在
	3.检测判断变量的类型
	4.检测判断是否使用了指定接口类型和3一样
	*/
	type CompanyInfo struct {
		Address string
		Name string
		MarketValue int "json:'market'"
	}
	var ch =make(chan interface{},10)
	var userInfo = make(map[string]interface{})
	userInfo["company_info"] = CompanyInfo{Address:"baidu xierqi",Name:"baidu",MarketValue:12345}
	userInfo["name"] = "jackmas"
	userInfo["age"]  = 110
	ch<-userInfo
	ch<-[]int{1,2,3,4,5,6}
	ch<-map[string]string{"Type":"Test"}
	close(ch)
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("chan status:","Chan Closed!")
			break
		}
		if newValue,ok:=v.(map[string]interface{});ok{   // 判断接口类型
			if _,ok:=newValue["company_info"];ok{
				for ck, cv := range newValue {
					fmt.Println(ck,"->",cv)
				}
			}
		}
	}
}

