package main

import (
	"fmt"
	"sort"
	"time"
)

type UserInfo struct {
	Uid      string `json:"uid"`
	UserName string `json:"user_name"`
	Sex      int    `json:"sex"`
}

func main() {
	planList:=make([]UserInfo,0)
	planList=append(planList,UserInfo{
		Sex: 10,
	})
	planList=append(planList,UserInfo{
		Sex: 7,
	})
	planList=append(planList,UserInfo{
		Sex: 11,
	})
	sort.SliceStable(planList, func(i, j int) bool {
		return planList[i].Sex < planList[j].Sex
	})
	fmt.Println(planList[0].Sex)
	return

	arr := []int{1, 2, 3, 4, 5, 6}
	t := 0
	for _, v := range arr {
		if v < t {
			t = v
		}
	}
	fmt.Println(t)
	return
	var user = make(map[string]*UserInfo)
	uid := "0001"
	user[uid] = &UserInfo{
		Uid:      uid,
		UserName: "jack",
		Sex:      1,
	}
	user[uid].UserName = "polly"
	fmt.Println(user[uid])

	//获取本地location
	toBeCharge := "0001-01-01 00:00:01"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)

}
