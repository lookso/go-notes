package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"unsafe"
)

func main() {
	//QrCode()
	Auth()
	//Awesome()
}

type ClientParams struct {
	Pid       string `json:"pid"`
	Sp        string `json:"sp"`
	DeviceId  string `json:"device_id"`
	ClientIp  string `json:"client_ip"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Uid       string `json:"uid"`
	Token     string `json:"token"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Ticket    string `json:"ticket"`
}

func QrCode() (qrcode *Response, err error) {
	//param := ClientParams{
	//	Pid:       PID,
	//	Sp:        "QM",
	//	DeviceId:  "94cd84cfdaacbc90",
	//	ClientIp:  "127.0.0.1",
	//	Timestamp: time.Now().Unix(),                    // 1586931183
	//	Nonce:      fmt.Sprintf("{\"timestamp\": %d}", time.Now().Unix()), // "{\"timestamp\": 1586931183}",
	//}
	param := ClientParams{
		Pid:       PID,
		Sp:        "KG",
		DeviceId:  "MAA118908000L",
		ClientIp:  "127.0.0.1",
		Timestamp: time.Now().Unix(),                    // 1586931183
		Nonce:      fmt.Sprintf("{\"timestamp\": %d}", time.Now().Unix()), // "{\"timestamp\": 1586931183}",
	}

	//param := ClientParams{
	//	Pid:       PID,
	//	Sp:        "KW",
	//	DeviceId:  "KTS17P480310",
	//	ClientIp:  "127.0.0.1",
	//	Timestamp: time.Now().Unix(),                    // 1586931183
	//	Nonce:      fmt.Sprintf("{\"timestamp\": %d}", time.Now().Unix()), // "{\"timestamp\": 1586931183}",
	//}
	url := "https://thirdsso.kugou.com/v2/user/qrcode/get"
	qrcode, err = TMEPost(url, param)
	if err != nil {
		fmt.Println("qrcode err", err)
		return nil, err
	}
	return qrcode, nil

}

func Auth() {

	//param := ClientParams{
	//	Pid:       PID,
	//	Sp:        "KW",
	//	DeviceId:  "KTS17P480310",
	//	ClientIp:  "127.0.0.1",
	//	Timestamp: time.Now().Unix(),//1586931183,
	//	Nonce:     fmt.Sprintf("{\"timestamp\": %d}", time.Now().Unix()),//"{\"timestamp\": 1586931183}",
	//	Ticket:    "1590476036#-M8Ek_eLBUH0bXNsrJ-r",
	//}
	//param := ClientParams{
	//	Pid:       PID,
	//	Sp:        "QM",
	//	DeviceId:  "94cd84cfdaacbc90",
	//	ClientIp:  "127.0.0.1",
	//	Timestamp: time.Now().Unix(),//1586931183,
	//	Nonce:     fmt.Sprintf("{\"timestamp\": %d}", time.Now().Unix()),//"{\"timestamp\": 1586931183}",
	//	Ticket:    "1590547734#-M8J155gyTpWzXY4ibVS",
	//}
	param := ClientParams{
		Pid:       PID,
		Sp:        "KG",
		DeviceId:  "MAA118908000L",
		ClientIp:  "127.0.0.1",
		Timestamp: time.Now().Unix(),//1586931183,
		Nonce:     fmt.Sprintf("{\"timestamp\": %d}", time.Now().Unix()),//"{\"timestamp\": 1586931183}",
		Ticket:    "be7d269ce0a22add7ba0f13e2f0c7b413154",
	}
	fmt.Println("param",param)
	url := "https://thirdsso.kugou.com/v2/user/qrcode/auth"
	userData, err := TMEPost(url, param)
	if err != nil {
		fmt.Println("userdata err", err)
	}
	fmt.Println("ud", userData.Data)
}

func Awesome() {
	param := ClientParams{
		Pid:       PID,
		Sp:        "QM",
		DeviceId:  "KTS17P480310",
		ClientIp:  "127.0.0.1",
		Timestamp: time.Now().Unix(),
		Nonce:     fmt.Sprintf("{\"timestamp\": %d}", time.Now().Unix()),
		Uid:       "7925907585188565679",
		Token:     "59aeba788eca995fbb954c2aef52d875233ed7b846d63e27648ee4410f767feb6a7e2c0e3613e9cad211f11043bdbec5",
		Page:      1,
		Size:      10,
	}
	url := "https://thirdsso.kugou.com/v2/playlist/awesome"
	// tem music
	aweData, err := TMEPost(url, param)
	if err != nil {
		fmt.Println("tempost err", err)
	}
	fmt.Println("aweData",aweData.Data)
}

var (
	PID  = "200031"
	PKey = "59508a47621c4a81b8710598d84202ac"
)

type Response struct {
	ErrorCode int         `json:"error_code"`
	ErrorMsg  string      `json:"error_msg"`
	Data      interface{} `json:"data"`
}

func TMEPost(url string, param ClientParams) (*Response, error) {

	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	//fmt.Println("signature", sign(Bytes2str(body), PKey))
	//fmt.Println("str_body", string(body))
	payload := strings.NewReader(string(body))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("signature", sign(Bytes2str(body), PKey))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	fmt.Println("body", string(body))

	resp := &Response{}
	if err := json.Unmarshal(body, resp); err != nil {
		fmt.Println("unmarshal err", err)
	}
	return resp, nil
}

func sign(body string, signKey string) string {
	return MD5(body + signKey)
}
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// Bytes2str : Bytes2str高效转换
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}