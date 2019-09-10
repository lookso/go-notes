package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	publicKey, err := ioutil.ReadFile("./public.pem")
	if err != nil {
		fmt.Println(err)
	}
	data := map[string]interface{}{
		"timestamp": time.Now().Unix(),
	}
	encryptData, _ := RsaEncrypt(publicKey, data)
	sign := base64.StdEncoding.EncodeToString(encryptData)
	fmt.Println(sign)
}

/*
 * @公钥加密
 * @params publicKey []byte 公钥
 * @params data map[string]interface{} 加密规则数据
 * @return []byte 解密后数据
 */
func RsaEncrypt(publicKey []byte, data map[string]interface{}) ([]byte, error) {
	if len(publicKey) == 0 {
		return nil, errors.New("public key empty")
	}
	if len(data) == 0 {
		return nil, errors.New("encrypt data empty")
	}
	outputData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("json Marshal fail")
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	// 加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, outputData)
}
