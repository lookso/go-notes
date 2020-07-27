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

var rsaPrivateKey = `-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAO1JXIjANU3mShpe
FGORAJ5/IYQT3Rna5ilsrDn8L4ujURriqEYufy0CoJnv+SqM6kscoM7gixd4E8jh
tH16hFy7Sl+4Lw7/OrdS874UdpNo0EmkafPgxDc8G7o1HnOKDmNarJgWfE1bpb9u
lEpHGuzAygtGbB6vSzC6Bk2PeRqhAgMBAAECgYAbiAuIgmSs6S9n58qN0uEzgqSs
4nEg8tkJrWY+RBhwJz7d07aajgKLgmaH9eP/H+J/XI778emxi2kgQa/jbze0IYUN
pSYzOid1gY4ctW9JPJyGQ1qpB23Asl6EQPDlJiMP+EyZOxAeO8GH/sQ1imVyl0jc
4Pys8k61E5tVu7HnAQJBAPgcesHgmtZSeBW1BqTo36/fWaUxNwpYtDd1o0b/ao2D
0j6KXtQswet84lKCJw31loK55PA+oBcpqzDCPwzZ+2kCQQD01MYMqo13DsXe/d2y
Np4t6BZhfyexKJQd25krn/mIUUNLYd6BbHRHohJK7X+EMFikSv5HT+dmC/7tTw3e
nFZ5AkA+EJG8sfzJpDOZ6oDQ+9gI3KxGIHuOQQZD4U+I0RfMcq9DKcXy+YdA6yqK
TTiLy2VtKidU2bWeVbQXLGAtTNIRAkEAmgeO/gMC8ydRJ4SRyIACLiF4iygjQZqk
7M/uYnrH05JEgxV0lfo0gaieV5NpiTGdYudnaFgF3baoUIZO1IGJCQJAU2UByCbz
m2yQ1kc6xUVNIjs9w2YkGlmQTn2c8xholeyTznIJEiQK1VaCcDy+6Pjz5rX4zHzx
KsP6+ka4UaefcQ==
-----END PRIVATE KEY-----`



var rsaPublickey = `-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAO1JXIjANU3mShpeFGORAJ5/IYQT3Rna5ilsrDn8L4ujURriqEYufy0C
oJnv+SqM6kscoM7gixd4E8jhtH16hFy7Sl+4Lw7/OrdS874UdpNo0EmkafPgxDc8
G7o1HnOKDmNarJgWfE1bpb9ulEpHGuzAygtGbB6vSzC6Bk2PeRqhAgMBAAE=
-----END RSA PUBLIC KEY-----`

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

//签名
//func RsaSignWithSha256(data []byte, keyBytes []byte) []byte {
//	h := sha256.New()
//	h.Write(data)
//	hashed := h.Sum(nil)
//	block, _ := pem.Decode(keyBytes)
//	if block == nil {
//		panic(errors.New("private key error"))
//	}
//	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
//	if err != nil {
//		fmt.Println("ParsePKCS8PrivateKey err", err)
//		panic(err)
//	}
//	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
//	if err != nil {
//		fmt.Printf("Error from signing: %s\n", err)
//		panic(err)
//	}
//	return signature
//}
//
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


//openssl genrsa -out rsa_private_key.pem 1024
//openssl pkcs8 -topk8 -inform PEM -in rsa_private_key.pem -outform PEM -nocrypt -out private_key.pem
//openssl rsa -in rsa_private_key.pem -RSAPublicKey_out -out rsa_public_key.pem