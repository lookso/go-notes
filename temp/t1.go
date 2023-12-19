package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
)

//X-Genie-DeviceId:mac
//X-Genie-Timestamp:1699844017
//X-Genie-Nonce:10086
//X-Genie-Sign:99b02ece847c22b5ce2f7601e69593b8
//X-Genie-AppId:200050
//X-Genie-Version:1.0.0
//X-Genie-Platform:mac
//X-Genie-Osversion:test
//Content-Type:application/json
//Authorization:bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwcm9kdWNlcl9pZCI6NTMwNDU1NTQ4LCJ0YWxfaWQiOiJUYWx6UmxyUS1ocXhySkt3YmI3ckpPNmVLdyIsImV4cCI6MTcwNDI3MjQ4Mn0.flAlNzSov6DlXr88VKUoZOmd78nAQmkVsfzMAwMPG3g
//Authorization:bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwcm9kdWNlcl9pZCI6MTM5MTA1MDgwNywidGFsX2lkIjoiVGFsenlHWndEdjdhTHZSWXIyMW9lRUFFd2ciLCJleHAiOjE3MDQwMTM5OTN9.U1luikYt-C3oqSEKKKIqwdYi2GxJ2xnNYOzB1vAQmAo
//
//
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
//  signConf.SignKey + "&X-Genie-Timestamp=" + timeStamp + "&X-Genie-Nonce=" + nonce + "&X-Genie-DeviceId=" + deviceId
func main()  {
	fmt.Println("rand.Perm(10)",rand.Perm(100))
	sandArr := rand.Perm(10)[:1]
	fmt.Println("sandArr", sandArr)
	template := 0
	for _, n := range sandArr {
		template |= 1 << n
	}
	fmt.Println(template)

	fmt.Println("8&(1<<(15+1-16))",8&(1<<(15+1-16)))
	if 8&(1<<(15+1-16)) > 0 {
		fmt.Println("true")
	}


	//fmt.Println(MD5("2c4a6975d9194d0cb279f75d230967f7&X-Genie-Timestamp=1702298029&X-Genie-Nonce=10086&X-Genie-DeviceId=mac"))
}


