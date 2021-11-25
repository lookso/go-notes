package main

import "fmt"

type UserInfo struct {
	Uid      string `json:"uid"`
	UserName string `json:"user_name"`
	Sex      int    `json:"sex"`
}

func main() {
	var user = make(map[string]*UserInfo)
	uid := "0001"
	user[uid] = &UserInfo{
		Uid:      uid,
		UserName: "jack",
		Sex:      1,
	}
	user[uid].UserName="polly"
	fmt.Println(user[uid])
}
