package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

func main() {

	fmt.Println("t", getT())
	fmt.Println("sign", getSign())
}

var (
	s1          = 1323692181
	s2          = "a23ng$m10Ab"
	currentTime = 1567482446
)

func getSign() string {
	return md5Func(strconv.FormatInt(time.Now().Unix(), 10) + s2 + "1.0" + "1.0")

}

func getT() (t int) {
	t = s1 ^ int(time.Now().Unix())
	return t
}
func md5Func(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
