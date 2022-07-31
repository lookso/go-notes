package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-notes/assembly/redis/common"
)

// GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
//127.0.0.1:6379> GEODIST geo_hash 天府广场 四川大剧院
//"213.2058"

func main() {
	ctx:=context.Background()
	client, _ := common.RedisClient()
	client.GeoAdd(ctx,"geo_hash", &redis.GeoLocation{
		Name:      "天府广场",
		Longitude: 104.072833,
		Latitude:  30.663422,
	}, &redis.GeoLocation{
		Name:      "四川大剧院",
		Longitude: 104.074378,
		Latitude:  30.664804,
	}, &redis.GeoLocation{
		Name:      "新华文轩",
		Longitude: 104.070084,
		Latitude:  30.664649,
	}, &redis.GeoLocation{
		Name:      "手工茶",
		Longitude: 104.072402,
		Latitude:  30.664121,
	}, &redis.GeoLocation{
		Name:      "宽窄巷子",
		Longitude: 104.059826,
		Latitude:  30.669883,
	}, &redis.GeoLocation{
		Name:      "奶茶",
		Longitude: 104.06085,
		Latitude:  30.670054,
	}, &redis.GeoLocation{
		Name:      "钓鱼台",
		Longitude: 104.058424,
		Latitude:  30.670737,
	}).Result()
	// GeoPos从key里返回所有给定位置元素的位置（经度和纬度）
	resPos, err := client.GeoPos(ctx,"geo_hash", "天府广场", "宽窄巷子").Result()
	if err != nil {
		fmt.Println(err)
	}
	//{104.07283455133438 30.663422957895442}
	//{104.05982583761215 30.66988396213059}
	for _, v := range resPos {
		fmt.Println("geoPos", *v)
	}
	// GeoHash返回一个或多个位置元素的 Geohash 表示
	resHash, err := client.GeoHash(ctx,"geo_hash", "天府广场", "四川大剧院").Result()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range resHash {
		fmt.Println("geoHash", v)
	}

	resRadiu, _ := client.GeoRadius(ctx,"geo_hash", 104.072833, 30.663422, &redis.GeoRadiusQuery{
		Radius:      800,   //radius表示范围距离，
		Unit:        "m",   //距离单位是 m|km|ft|mi
		WithCoord:   true,  //传入WITHCOORD参数，则返回结果会带上匹配位置的经纬度
		WithDist:    true,  //传入WITHDIST参数，则返回结果会带上匹配位置与给定地理位置的距离。
		WithGeoHash: true,  //传入WITHHASH参数，则返回结果会带上匹配位置的hash值。
		Count:       4,     //入COUNT参数，可以返回指定数量的结果。
		Sort:        "ASC", //默认结果是未排序的，传入ASC为从近到远排序，传入DESC为从远到近排序。
	}).Result()
	for _, v := range resRadiu {
		fmt.Println("resRadiu", v)
	}

}
