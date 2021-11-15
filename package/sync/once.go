/*
@Time : 2019/5/11 10:54 AM
@Author : Tenlu
@File : once
@Software: GoLand
*/
package main

import (
	"sync"
)

type ScrmOpen struct {
	appId  string
	appKey string
}
type ScrmOpen2 struct {
	appId  string
	appKey string
}

func main() {
	//fmt.Println(NewScrmOpenApi(1))
	//fmt.Println(NewScrmOpenApi(1))
	//fmt.Println(NewScrmOpenApi(2))
	//fmt.Println(NewScrmOpenApi(2))
	//fmt.Println(NewScrmOpenApi(1))
	NewScrmOpenApi(1)

	NewScrmOpenApi(2)

}

//var once sync.Once
// 这个地方有个坑,只在外部声明一个统一的once,两个实例方法都使用这个once的时候,就会发现只有一个once被实例化成功,另一个实例会报未初始化,
// 所以需要在子方法内容单独声明 sync.once
type ScrmOpenApiV2 struct {
	appId  string `json:"app_id"`
}

var instance *ScrmOpenApiV2
var instance2 *ScrmOpenApiV2

func NewScrmOpenApi(t int) *ScrmOpenApiV2 {
	if t == 1 {
		return onceFunc()
	}
	return newOnceFunc()
}
func onceFunc() *ScrmOpenApiV2 {
	if instance != nil {
		return instance
	}
	var once sync.Once
	once.Do(func() {
		instance = new(ScrmOpenApiV2)
		instance.appId = "100"
	})
	return instance
}
func newOnceFunc() *ScrmOpenApiV2 {
	var once sync.Once
	once.Do(func() {
		instance2 = new(ScrmOpenApiV2)
		instance2.appId = "200"
	})
	return instance2
}

