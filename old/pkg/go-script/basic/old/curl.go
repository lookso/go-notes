package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://local-api-sns-template-healthcare.qschou.com/api/sns/sns_3_null_2_community/get?id=community_illness"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("Cookie", "qsc-token=4cfe83d90c8a25a17db36e61273ac96da2738f62;user_id=1679335")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "0cbe3fc0-9ba2-4c1f-9987-c5a26d2a0d1e")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
