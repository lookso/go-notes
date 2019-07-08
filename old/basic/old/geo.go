package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	GetNetIP()
}

var getIp = flag.String("get_ip", "", "external|internal")

func GetNetIP() {
	fmt.Println("Usage of ./getmyip --get_ip=(external|internal)")
	flag.Parse()
	if *getIp == "external" {
		GetExternal()
	}
	if *getIp == "internal" {
		GetInternal()
	}
}

func GetExternal() {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	os.Exit(0)
}

func GetInternal() {
	adders, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}

	for _, a := range adders {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
	os.Exit(0)
}
