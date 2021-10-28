/*
@Time : 2019/5/11 10:54 AM
@Author : Tenlu
@File : once
@Software: GoLand
*/
package main

import (
	"fmt"
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

	fmt.Println(NewScrmOpenApi(2))
	fmt.Println(NewScrmOpenApi(1))
	fmt.Println(NewScrmOpenApi(1))
	fmt.Println(NewScrmOpenApi(2))
}

var once sync.Once

type ScrmOpenApiV2 struct {
	AppId string `json:"app_id"`
}

var instance *ScrmOpenApiV2
var instance2 *ScrmOpenApiV2

func NewScrmOpenApi(t int) *ScrmOpenApiV2 {
	if t==2{
		return once2()
	}
	return newonce()
}
func once2() *ScrmOpenApiV2 {
	if instance != nil {
		return instance
	}
	var once sync.Once
	once.Do(func() {
		instance=new(ScrmOpenApiV2)
		instance.AppId="1"
	})
	return instance
}
func newonce()  *ScrmOpenApiV2{
	var once sync.Once
	once.Do(func() {
		instance2=new(ScrmOpenApiV2)
		instance2.AppId="2"
	})
	return instance2
}
