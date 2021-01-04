/*
@Time : 2019/6/2 8:51 PM 
@Author : Tenlu
@File : genkey
@Software: GoLand
*/
package main

import (
	"crypto/x509"
	"encoding/pem"
	"os"
	"flag"
	"crypto/rsa"
	"crypto/rand"
	"log"
)

// 由genkey.go生成公钥和私文件，在rsa.go里使用生成的公钥和私钥进行加密和解密
func main() {
	var bits int
	flag.IntVar(&bits, "b", 1024, "秘钥长度，默认为1024")
	if err := GenRsaKey(bits); err != nil {
		log.Fatal("秘钥文件生成失败")
	}
	log.Println("秘钥文件生成成功")
}

//生成 私钥和公钥文件
func GenRsaKey(bits int) error {
	//生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	//生成公钥文件
	publicKey := &privateKey.PublicKey
	defPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: defPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

