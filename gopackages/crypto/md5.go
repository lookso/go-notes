/*
@Time : 2019/5/25 10:56 PM
@Author : Tenlu
@File : md5
@Software: GoLand
*/
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"io"
)

func main() {
	//timeStamp := cast.ToString(time.Now().Unix())
	//fmt.Printf("timeStamp:%s\n", timeStamp)
	//nonce := uuid.New().String()
	nonce:="12321323"
	deviceId := "spider"
	timeStamp := "1698823216"
	signStr := "f19ff4ac8cb0491c8538c89418315072&X-Genie-Timestamp=" + timeStamp + "&X-Genie-Nonce=" + nonce + "&X-Genie-DeviceId=" + deviceId
	fmt.Println(GetMd5(signStr))

	buildSignHeader("100001","1qw9eq8we2cd4ef35try1rt1y7vfg5tr9")
}
//
//
//key: "100001"
//secret: "1qw9eq8we2cd4ef35try1rt1y7vfg5tr9"


//"200052": # 众包管理后台
//secret: ""
//url: ""
//type: 2 #1:url鉴权 2:secret鉴权 3:不鉴权
//sign_key: "f19ff4ac8cb0491c8538c89418315072"
//application: "zb_back"


func buildSignHeader(key, secret string) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := uuid.New().String()
	deviceId := "123456"
	fmt.Println("timestamp",timestamp)

	sum := md5.Sum([]byte("f19ff4ac8cb0491c8538c89418315072&X-Genie-Timestamp=1700050444"  + "&X-Genie-Nonce=1234"  + "&X-Genie-DeviceId=123456"))
	sign := hex.EncodeToString(sum[:])
	fmt.Println("sign",sign)

	return map[string]string{
		"Content-Type":      "application/json",
		"X-Genie-TraceId":   uuid.New().String(),
		"X-Genie-DeviceId":  deviceId,
		"X-Genie-Timestamp": timestamp,
		"X-Genie-Nonce":     nonce,
		"X-Genie-AppId":     key,
		"X-Genie-Sign":      sign,
		"X-Genie-Version":   "V230223",
		"X-Genie-Osversion": "1.0.0",
		"X-Genie-Platform":  "web",
	}
}



func GetMd5(str string) string {
	h := md5.New()
	io.WriteString(h, str)

	return fmt.Sprintf("%x", h.Sum(nil))
}


// secret: "PDA749CABC9DE50FF57503BC8123143D2B"



//sign_str = "PDQueryAsDfGhJkLZxCv&X-Genie-Timestamp=1698647044&X-Genie-Nonce=XfkMLkTLtEzWhVZc&X-Genie-DeviceId=spider"
//
//headers = {'X-Genie-DeviceId': 'spider', 'X-Genie-Timestamp': 1698647044, 'X-Genie-Nonce': 'XfkMLkTLtEzWhVZc', 'X-Genie-AppId': 300000, 'X-Genie-Version': '1.0.0', 'X-Genie-Platform': 'spider-server', 'X-Genie-Sign': '8f5f983fcccc605134df477c7968df23'}
