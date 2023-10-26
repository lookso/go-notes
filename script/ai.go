package main

//import (
//	"crypto/hmac"
//	"crypto/sha1"
//	"encoding/base64"
//	"encoding/json"
//	"fmt"
//	"github.com/google/uuid"
//	"io/ioutil"
//	"net/http"
//	netUrl "net/url"
//	"sort"
//	"strings"
//	"time"
//)

// var imageUrl="https://qz-test.oss-cn-beijing.aliyuncs.com/07/KET%E4%BD%9C%E6%96%87-1/%E7%9C%8B%E5%9B%BE%E5%86%99%E6%95%85%E4%BA%8B.jpg"

//func main() {
//	url := "http://openai.100tal.com/aitext/english-composition/text-correction"
//	urlParams := make(map[string]string)
//	bodyParams := make(map[string]interface{})
//	bodyParams["image_url"] = []string{"https://qz-test.oss-cn-beijing.aliyuncs.com/07/KET%E4%BD%9C%E6%96%87-1/%E7%9C%8B%E5%9B%BE%E5%86%99%E6%95%85%E4%BA%8B.jpg"}
//	accessKeyID := "920701612630999040"
//	accessKeySecret := "f4f5fd3f1e4b46618e9d67676d2eb288"
//
//	run(url, urlParams, bodyParams, accessKeyID, accessKeySecret)
//}
//
//func run(reqUrl string, urlParams map[string]string, bodyParams map[string]interface{}, accessKeyID string, accessKeySecret string) {
//	headers := make(http.Header)
//	headers.Set("Content-Type", "application/json")
//
//	urlParams["access_key_id"] = accessKeyID
//	urlParams["timestamp"] = time.Now().Format("2006-01-02T15:04:05")
//	urlParams["signature_nonce"] = generateSignatureNonce()
//
//	signParams := make(map[string]interface{})
//	bodyJson, err := json.Marshal(bodyParams)
//	signParams["request_body"] = string(bodyJson)
//	for key, value := range urlParams {
//		signParams[key] = value
//	}
//
//	signature := generateSignature(signParams, accessKeySecret)
//
//	urlParams["signature"] = netUrl.QueryEscape(signature)
//
//	urlWithParams := reqUrl + "?" + urlFormat(urlParams)
//
//	body, err := json.Marshal(bodyParams)
//	if err != nil {
//		fmt.Println("Failed to marshal bodyParams:", err)
//		return
//	}
//
//	fmt.Println("body", string(body))
//	req, err := http.NewRequest(http.MethodPost, urlWithParams, strings.NewReader(string(body)))
//	if err != nil {
//		fmt.Println("Failed to create request:", err)
//		return
//	}
//	req.Header = headers
//
//	client := http.DefaultClient
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println("Failed to send request:", err)
//		return
//	}
//	defer resp.Body.Close()
//
//	responseBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Failed to read response body:", err)
//		return
//	}
//
//	fmt.Println(string(responseBody))
//}
//
//func generateSignatureNonce() string {
//	return uuid.New().String()
//}
//
//func generateSignature(params map[string]interface{}, accessKeySecret string) string {
//	keys := make([]string, 0, len(params))
//
//	for key := range params {
//		keys = append(keys, key)
//	}
//	sort.Strings(keys)
//
//	paramList := make([]string, 0, len(params))
//	for _, key := range keys {
//		value := fmt.Sprintf("%v", params[key])
//		paramList = append(paramList, fmt.Sprintf("%s=%s", key, value))
//	}
//	stringToSign := strings.Join(paramList, "&")
//	fmt.Println("stringToSign",stringToSign)
//	secret := accessKeySecret + "&"
//
//	h := hmac.New(sha1.New, []byte(secret))
//	h.Write([]byte(stringToSign))
//	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
//
//	return signature
//}
//
//func urlFormat(params map[string]string) string {
//	keys := make([]string, 0, len(params))
//	for key := range params {
//		keys = append(keys, key)
//	}
//	sort.Strings(keys)
//	fmt.Println("sort_keys", keys)
//
//	paramList := make([]string, 0, len(params))
//	for _, key := range keys {
//		value := params[key]
//		paramList = append(paramList, fmt.Sprintf("%s=%s", key, netUrl.QueryEscape(value)))
//	}
//
//	return strings.Join(paramList, "&")
//}
