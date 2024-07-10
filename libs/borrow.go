package main

import (
	"log"

	"github.com/pkg/browser"
	"github.com/pkg/profile"
)

func main() {
	// 开发测试阶段,启动CPU分析，程序结束时自动停止分析,类似pprof
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	url := "https://www.baidu.com"
	err := browser.OpenURL(url)
	if err != nil {
		log.Fatalf("failed to open URL: %v", err)
	}
}
