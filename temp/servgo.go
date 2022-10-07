package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type SearchReq struct {
	PageNum          int             `json:"pageNum"`
	PageSize         int             `json:"pageSize"`
	FilterID         int             `json:"filterId"`
	ObjectAPIName    string          `json:"objectApiName"`
	JudgeStrategy    int             `json:"judgeStrategy"`
	CustomJudgeLogic string          `json:"customJudgeLogic"`
	ConditionList    []ConditionList `json:"conditionList"`
}

type ConditionList struct {
	FieldAPIName string `json:"fieldApiName"`
	Operator     string `json:"operator"`
	Value        string `json:"value"`
}

func main() {
	email := "pepg@baidu.com"
	apiToken := "526cfde851f0b723f7f91955e98781f6"
	timestamp := time.Now().Unix()
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s&%s&%d", email, apiToken, timestamp)))
	sign := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("sign", sign)
	url := fmt.Sprintf("https://servicego.udesk.cn/api/v1/data?email=%s&timestamp=%d&sign=%s&", email, timestamp, sign)

	method := "PUT"
	payload := strings.NewReader(`{"objectApiName":"rpatemplate","uniqueFieldApiName":"addid","uniqueFieldValue":"1317","fieldDataList":[{"fieldApiName":"s4id","fieldValue":"test_85","foreignExternalFieldApiName":""}]}`)
	//{
	//	"objectApiName": "rpatemplate",
	//	"id": 0,
	//	"uniqueFieldApiName": "addid",
	//	"uniqueFieldValue": "1317",
	//	"fieldDataList": [
	//{
	//"fieldApiName": "s4id",
	//"fieldValue": "test_85",
	//"foreignExternalFieldApiName": ""
	//}
	//]
	//}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
