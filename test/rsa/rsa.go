package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

var rsaPrivateKey=`xxx`

var rsaPublickey =`xxx`


func main()  {
	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	var data = `{"pid":"200021","sp":"KG","device_id":"MAA118908000L","client_ip":"127.0.0.1"}`
	fmt.Println("对消息进行签名操作...")
	signData,_ := RsaSignWithSha256([]byte(data), []byte(rsaPrivateKey))
	sg := base64.StdEncoding.EncodeToString(signData)
	fmt.Println("sg",sg)
	//fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	fmt.Println("对消息进行验签操作...")
	if RsaVerySignWithSha256([]byte(data), sg, []byte(rsaPublickey)) {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}
}
// 签名
func RsaSignWithSha256(data []byte, keyBytes []byte) (sign []byte, err error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil,errors.New("private key error")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		return nil,err
	}
	priKey := privateKey.(*rsa.PrivateKey)
	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return nil,err
	}
	return signature,nil
}

//验签
func RsaVerySignWithSha256(data []byte, sg string, keyBytes []byte) bool {
	signData,_ := base64.StdEncoding.DecodeString(sg)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}