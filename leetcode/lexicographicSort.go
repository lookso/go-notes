/*
@Time : 2019/5/25 9:58 PM 
@Author : Tenlu
@File : lexicographicSort
@Software: GoLand
*/
package main

import (
	"sort"
	"crypto/sha1"
	"encoding/hex"
	"strings"
	"strconv"
	"time"
	"fmt"
	"math/rand"
)
// 字典排序

const Token = "d6a0b3aacfe5417d60q5s23f3e21ac21"
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init(){
	//以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}
func main(){
	Nonce := randSeq(10)
	signature,_:=CreateSign(Token,strconv.Itoa(int(time.Now().Unix())),Nonce)
	fmt.Println("Signature:",signature)
}
// 按字典序生成签名
func CreateSign(Token string,TimeStamp string,Nonce string) (string,error) {
	var signature string
	strArr:=[]string{Token,TimeStamp,Nonce}
	sort.Strings(strArr)
	var content string
	content = ""
	for _,v:=range strArr{
		content += v
	}
	hash := sha1.New()
	hash.Write([]byte(content))
	cipherStr:=hash.Sum(nil)
	// 二进制数组转换为16进制
	signature = strings.ToUpper(hex.EncodeToString(cipherStr))
	return signature,nil
}

// 返回指定长度得字符串随机数
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

